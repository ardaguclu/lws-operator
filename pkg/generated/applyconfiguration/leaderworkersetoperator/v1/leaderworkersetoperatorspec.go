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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apioperatorv1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// LeaderWorkerSetOperatorSpecApplyConfiguration represents a declarative configuration of the LeaderWorkerSetOperatorSpec type for use
// with apply.
type LeaderWorkerSetOperatorSpecApplyConfiguration struct {
	operatorv1.OperatorSpecApplyConfiguration `json:",inline"`
}

// LeaderWorkerSetOperatorSpecApplyConfiguration constructs a declarative configuration of the LeaderWorkerSetOperatorSpec type for use with
// apply.
func LeaderWorkerSetOperatorSpec() *LeaderWorkerSetOperatorSpecApplyConfiguration {
	return &LeaderWorkerSetOperatorSpecApplyConfiguration{}
}

// WithManagementState sets the ManagementState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagementState field is set to the value of the last call.
func (b *LeaderWorkerSetOperatorSpecApplyConfiguration) WithManagementState(value apioperatorv1.ManagementState) *LeaderWorkerSetOperatorSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.ManagementState = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *LeaderWorkerSetOperatorSpecApplyConfiguration) WithLogLevel(value apioperatorv1.LogLevel) *LeaderWorkerSetOperatorSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.LogLevel = &value
	return b
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *LeaderWorkerSetOperatorSpecApplyConfiguration) WithOperatorLogLevel(value apioperatorv1.LogLevel) *LeaderWorkerSetOperatorSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.OperatorLogLevel = &value
	return b
}

// WithUnsupportedConfigOverrides sets the UnsupportedConfigOverrides field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnsupportedConfigOverrides field is set to the value of the last call.
func (b *LeaderWorkerSetOperatorSpecApplyConfiguration) WithUnsupportedConfigOverrides(value runtime.RawExtension) *LeaderWorkerSetOperatorSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.UnsupportedConfigOverrides = &value
	return b
}

// WithObservedConfig sets the ObservedConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedConfig field is set to the value of the last call.
func (b *LeaderWorkerSetOperatorSpecApplyConfiguration) WithObservedConfig(value runtime.RawExtension) *LeaderWorkerSetOperatorSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.ObservedConfig = &value
	return b
}
