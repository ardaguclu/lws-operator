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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	leaderworkersetoperatorv1 "github.com/openshift/lws-operator/pkg/apis/leaderworkersetoperator/v1"
	applyconfigurationleaderworkersetoperatorv1 "github.com/openshift/lws-operator/pkg/generated/applyconfiguration/leaderworkersetoperator/v1"
	scheme "github.com/openshift/lws-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// LeaderWorkerSetOperatorsGetter has a method to return a LeaderWorkerSetOperatorInterface.
// A group's client should implement this interface.
type LeaderWorkerSetOperatorsGetter interface {
	LeaderWorkerSetOperators() LeaderWorkerSetOperatorInterface
}

// LeaderWorkerSetOperatorInterface has methods to work with LeaderWorkerSetOperator resources.
type LeaderWorkerSetOperatorInterface interface {
	Create(ctx context.Context, leaderWorkerSetOperator *leaderworkersetoperatorv1.LeaderWorkerSetOperator, opts metav1.CreateOptions) (*leaderworkersetoperatorv1.LeaderWorkerSetOperator, error)
	Update(ctx context.Context, leaderWorkerSetOperator *leaderworkersetoperatorv1.LeaderWorkerSetOperator, opts metav1.UpdateOptions) (*leaderworkersetoperatorv1.LeaderWorkerSetOperator, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, leaderWorkerSetOperator *leaderworkersetoperatorv1.LeaderWorkerSetOperator, opts metav1.UpdateOptions) (*leaderworkersetoperatorv1.LeaderWorkerSetOperator, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*leaderworkersetoperatorv1.LeaderWorkerSetOperator, error)
	List(ctx context.Context, opts metav1.ListOptions) (*leaderworkersetoperatorv1.LeaderWorkerSetOperatorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *leaderworkersetoperatorv1.LeaderWorkerSetOperator, err error)
	Apply(ctx context.Context, leaderWorkerSetOperator *applyconfigurationleaderworkersetoperatorv1.LeaderWorkerSetOperatorApplyConfiguration, opts metav1.ApplyOptions) (result *leaderworkersetoperatorv1.LeaderWorkerSetOperator, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, leaderWorkerSetOperator *applyconfigurationleaderworkersetoperatorv1.LeaderWorkerSetOperatorApplyConfiguration, opts metav1.ApplyOptions) (result *leaderworkersetoperatorv1.LeaderWorkerSetOperator, err error)
	LeaderWorkerSetOperatorExpansion
}

// leaderWorkerSetOperators implements LeaderWorkerSetOperatorInterface
type leaderWorkerSetOperators struct {
	*gentype.ClientWithListAndApply[*leaderworkersetoperatorv1.LeaderWorkerSetOperator, *leaderworkersetoperatorv1.LeaderWorkerSetOperatorList, *applyconfigurationleaderworkersetoperatorv1.LeaderWorkerSetOperatorApplyConfiguration]
}

// newLeaderWorkerSetOperators returns a LeaderWorkerSetOperators
func newLeaderWorkerSetOperators(c *OpenShiftOperatorV1Client) *leaderWorkerSetOperators {
	return &leaderWorkerSetOperators{
		gentype.NewClientWithListAndApply[*leaderworkersetoperatorv1.LeaderWorkerSetOperator, *leaderworkersetoperatorv1.LeaderWorkerSetOperatorList, *applyconfigurationleaderworkersetoperatorv1.LeaderWorkerSetOperatorApplyConfiguration](
			"leaderworkersetoperators",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *leaderworkersetoperatorv1.LeaderWorkerSetOperator {
				return &leaderworkersetoperatorv1.LeaderWorkerSetOperator{}
			},
			func() *leaderworkersetoperatorv1.LeaderWorkerSetOperatorList {
				return &leaderworkersetoperatorv1.LeaderWorkerSetOperatorList{}
			},
		),
	}
}
