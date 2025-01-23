package operator

import (
	"context"
	"os"
	"time"

	apiextclientv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/openshift/library-go/pkg/operator/loglevel"

	operatorconfigclient "github.com/openshift/lws-operator/pkg/generated/clientset/versioned"
	operatorclientinformers "github.com/openshift/lws-operator/pkg/generated/informers/externalversions"
	"github.com/openshift/lws-operator/pkg/operator/operatorclient"
)

const (
	podNamespaceEnv   = "POD_NAMESPACE"
	operatorNamespace = "openshift-lws-operator"
	workQueueKey      = "key"
)

func RunOperator(ctx context.Context, cc *controllercmd.ControllerContext) error {
	kubeClient, err := kubernetes.NewForConfig(cc.ProtoKubeConfig)
	if err != nil {
		return err
	}

	dynamicClient, err := dynamic.NewForConfig(cc.ProtoKubeConfig)
	if err != nil {
		return err
	}

	apiextensionClient, err := apiextclientv1.NewForConfig(cc.KubeConfig)
	if err != nil {
		return err
	}

	operatorConfigClient, err := operatorconfigclient.NewForConfig(cc.KubeConfig)
	if err != nil {
		return err
	}
	operatorConfigInformers := operatorclientinformers.NewSharedInformerFactory(operatorConfigClient, 10*time.Minute)

	namespace := getNamespace()

	lwsOperatorClient := &operatorclient.LWSOperatorClient{
		Ctx:               ctx,
		SharedInformer:    operatorConfigInformers.LwsOperators().V1alpha1().LwsOperators().Informer(),
		OperatorClient:    operatorConfigClient.LwsOperatorsV1alpha1(),
		OperatorNamespace: namespace,
	}

	targetConfigReconciler := NewTargetConfigReconciler(
		ctx,
		os.Getenv("RELATED_IMAGE_OPERAND_IMAGE"),
		namespace,
		operatorConfigClient.LwsOperatorsV1alpha1(),
		operatorConfigInformers.LwsOperators().V1alpha1().LwsOperators(),
		lwsOperatorClient,
		dynamicClient,
		kubeClient,
		apiextensionClient,
		cc.EventRecorder,
	)

	logLevelController := loglevel.NewClusterOperatorLoggingController(lwsOperatorClient, cc.EventRecorder)

	klog.Infof("Starting informers")
	operatorConfigInformers.Start(ctx.Done())

	klog.Infof("Starting log level controller")
	go logLevelController.Run(ctx, 1)
	klog.Infof("Starting target config reconciler")
	go targetConfigReconciler.Run(1, ctx.Done())

	<-ctx.Done()
	return nil
}

// getNamespace returns in-cluster namespace
func getNamespace() string {
	if nsBytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		return string(nsBytes)
	}
	if podNamespace := os.Getenv(podNamespaceEnv); len(podNamespace) > 0 {
		return podNamespace
	}
	return operatorNamespace
}
