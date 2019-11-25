// +build !ignore_autogenerated

/*
Copyright 2019 Giant Swarm GmbH.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSCluster) DeepCopyInto(out *AWSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSCluster.
func (in *AWSCluster) DeepCopy() *AWSCluster {
	if in == nil {
		return nil
	}
	out := new(AWSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterList) DeepCopyInto(out *AWSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterList.
func (in *AWSClusterList) DeepCopy() *AWSClusterList {
	if in == nil {
		return nil
	}
	out := new(AWSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpec) DeepCopyInto(out *AWSClusterSpec) {
	*out = *in
	out.Cluster = in.Cluster
	out.Provider = in.Provider
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpec.
func (in *AWSClusterSpec) DeepCopy() *AWSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecCluster) DeepCopyInto(out *AWSClusterSpecCluster) {
	*out = *in
	out.DNS = in.DNS
	out.OIDC = in.OIDC
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecCluster.
func (in *AWSClusterSpecCluster) DeepCopy() *AWSClusterSpecCluster {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecClusterDNS) DeepCopyInto(out *AWSClusterSpecClusterDNS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecClusterDNS.
func (in *AWSClusterSpecClusterDNS) DeepCopy() *AWSClusterSpecClusterDNS {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecClusterDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecClusterOIDC) DeepCopyInto(out *AWSClusterSpecClusterOIDC) {
	*out = *in
	out.Claims = in.Claims
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecClusterOIDC.
func (in *AWSClusterSpecClusterOIDC) DeepCopy() *AWSClusterSpecClusterOIDC {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecClusterOIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecClusterOIDCClaims) DeepCopyInto(out *AWSClusterSpecClusterOIDCClaims) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecClusterOIDCClaims.
func (in *AWSClusterSpecClusterOIDCClaims) DeepCopy() *AWSClusterSpecClusterOIDCClaims {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecClusterOIDCClaims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecProvider) DeepCopyInto(out *AWSClusterSpecProvider) {
	*out = *in
	out.CredentialSecret = in.CredentialSecret
	out.Master = in.Master
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecProvider.
func (in *AWSClusterSpecProvider) DeepCopy() *AWSClusterSpecProvider {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecProviderCredentialSecret) DeepCopyInto(out *AWSClusterSpecProviderCredentialSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecProviderCredentialSecret.
func (in *AWSClusterSpecProviderCredentialSecret) DeepCopy() *AWSClusterSpecProviderCredentialSecret {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecProviderCredentialSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecProviderMaster) DeepCopyInto(out *AWSClusterSpecProviderMaster) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecProviderMaster.
func (in *AWSClusterSpecProviderMaster) DeepCopy() *AWSClusterSpecProviderMaster {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecProviderMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterStatus) DeepCopyInto(out *AWSClusterStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Cluster.DeepCopyInto(&out.Cluster)
	out.Provider = in.Provider
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterStatus.
func (in *AWSClusterStatus) DeepCopy() *AWSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AWSClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterStatusProvider) DeepCopyInto(out *AWSClusterStatusProvider) {
	*out = *in
	out.Network = in.Network
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterStatusProvider.
func (in *AWSClusterStatusProvider) DeepCopy() *AWSClusterStatusProvider {
	if in == nil {
		return nil
	}
	out := new(AWSClusterStatusProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterStatusProviderNetwork) DeepCopyInto(out *AWSClusterStatusProviderNetwork) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterStatusProviderNetwork.
func (in *AWSClusterStatusProviderNetwork) DeepCopy() *AWSClusterStatusProviderNetwork {
	if in == nil {
		return nil
	}
	out := new(AWSClusterStatusProviderNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeployment) DeepCopyInto(out *AWSMachineDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeployment.
func (in *AWSMachineDeployment) DeepCopy() *AWSMachineDeployment {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSMachineDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentList) DeepCopyInto(out *AWSMachineDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSMachineDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentList.
func (in *AWSMachineDeploymentList) DeepCopy() *AWSMachineDeploymentList {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSMachineDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentSpec) DeepCopyInto(out *AWSMachineDeploymentSpec) {
	*out = *in
	out.NodePool = in.NodePool
	in.Provider.DeepCopyInto(&out.Provider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentSpec.
func (in *AWSMachineDeploymentSpec) DeepCopy() *AWSMachineDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentSpecNodePool) DeepCopyInto(out *AWSMachineDeploymentSpecNodePool) {
	*out = *in
	out.Machine = in.Machine
	out.Scaling = in.Scaling
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentSpecNodePool.
func (in *AWSMachineDeploymentSpecNodePool) DeepCopy() *AWSMachineDeploymentSpecNodePool {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentSpecNodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentSpecNodePoolMachine) DeepCopyInto(out *AWSMachineDeploymentSpecNodePoolMachine) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentSpecNodePoolMachine.
func (in *AWSMachineDeploymentSpecNodePoolMachine) DeepCopy() *AWSMachineDeploymentSpecNodePoolMachine {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentSpecNodePoolMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentSpecNodePoolScaling) DeepCopyInto(out *AWSMachineDeploymentSpecNodePoolScaling) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentSpecNodePoolScaling.
func (in *AWSMachineDeploymentSpecNodePoolScaling) DeepCopy() *AWSMachineDeploymentSpecNodePoolScaling {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentSpecNodePoolScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentSpecProvider) DeepCopyInto(out *AWSMachineDeploymentSpecProvider) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Worker = in.Worker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentSpecProvider.
func (in *AWSMachineDeploymentSpecProvider) DeepCopy() *AWSMachineDeploymentSpecProvider {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentSpecProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentSpecProviderWorker) DeepCopyInto(out *AWSMachineDeploymentSpecProviderWorker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentSpecProviderWorker.
func (in *AWSMachineDeploymentSpecProviderWorker) DeepCopy() *AWSMachineDeploymentSpecProviderWorker {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentSpecProviderWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonClusterStatus) DeepCopyInto(out *CommonClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CommonClusterStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]CommonClusterStatusVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonClusterStatus.
func (in *CommonClusterStatus) DeepCopy() *CommonClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CommonClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonClusterStatusCondition) DeepCopyInto(out *CommonClusterStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonClusterStatusCondition.
func (in *CommonClusterStatusCondition) DeepCopy() *CommonClusterStatusCondition {
	if in == nil {
		return nil
	}
	out := new(CommonClusterStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonClusterStatusVersion) DeepCopyInto(out *CommonClusterStatusVersion) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonClusterStatusVersion.
func (in *CommonClusterStatusVersion) DeepCopy() *CommonClusterStatusVersion {
	if in == nil {
		return nil
	}
	out := new(CommonClusterStatusVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeepCopyTime.
func (in *DeepCopyTime) DeepCopy() *DeepCopyTime {
	if in == nil {
		return nil
	}
	out := new(DeepCopyTime)
	in.DeepCopyInto(out)
	return out
}
