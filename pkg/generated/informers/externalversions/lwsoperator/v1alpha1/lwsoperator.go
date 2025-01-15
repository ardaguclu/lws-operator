/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	apislwsoperatorv1alpha1 "github.com/openshift/lws-operator/pkg/apis/lwsoperator/v1alpha1"
	time "time"

	versioned "github.com/openshift/lws-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/lws-operator/pkg/generated/informers/externalversions/internalinterfaces"
	lwsoperatorv1alpha1 "github.com/openshift/lws-operator/pkg/generated/listers/lwsoperator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LwsOperatorInformer provides access to a shared informer and lister for
// LwsOperators.
type LwsOperatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() lwsoperatorv1alpha1.LwsOperatorLister
}

type lwsOperatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLwsOperatorInformer constructs a new informer for LwsOperator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLwsOperatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLwsOperatorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLwsOperatorInformer constructs a new informer for LwsOperator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLwsOperatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LwsOperatorsV1alpha1().LwsOperators(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LwsOperatorsV1alpha1().LwsOperators(namespace).Watch(context.TODO(), options)
			},
		},
		&apislwsoperatorv1alpha1.LwsOperator{},
		resyncPeriod,
		indexers,
	)
}

func (f *lwsOperatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLwsOperatorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *lwsOperatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apislwsoperatorv1alpha1.LwsOperator{}, f.defaultInformer)
}

func (f *lwsOperatorInformer) Lister() lwsoperatorv1alpha1.LwsOperatorLister {
	return lwsoperatorv1alpha1.NewLwsOperatorLister(f.Informer().GetIndexer())
}
