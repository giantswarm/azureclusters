//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Giant Swarm GmbH.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSCluster) DeepCopyInto(out *AWSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
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
	in.Provider.DeepCopyInto(&out.Provider)
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
	out.KubeProxy = in.KubeProxy
	out.OIDC = in.OIDC
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
func (in *AWSClusterSpecClusterKubeProxy) DeepCopyInto(out *AWSClusterSpecClusterKubeProxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecClusterKubeProxy.
func (in *AWSClusterSpecClusterKubeProxy) DeepCopy() *AWSClusterSpecClusterKubeProxy {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecClusterKubeProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecClusterOIDC) DeepCopyInto(out *AWSClusterSpecClusterOIDC) {
	*out = *in
	out.Claims = in.Claims
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
	out.Nodes = in.Nodes
	in.Pods.DeepCopyInto(&out.Pods)
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
func (in *AWSClusterSpecProviderNodes) DeepCopyInto(out *AWSClusterSpecProviderNodes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecProviderNodes.
func (in *AWSClusterSpecProviderNodes) DeepCopy() *AWSClusterSpecProviderNodes {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecProviderNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpecProviderPods) DeepCopyInto(out *AWSClusterSpecProviderPods) {
	*out = *in
	if in.ExternalSNAT != nil {
		in, out := &in.ExternalSNAT, &out.ExternalSNAT
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpecProviderPods.
func (in *AWSClusterSpecProviderPods) DeepCopy() *AWSClusterSpecProviderPods {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpecProviderPods)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterStatus) DeepCopyInto(out *AWSClusterStatus) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	out.Provider = in.Provider
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
func (in *AWSControlPlane) DeepCopyInto(out *AWSControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSControlPlane.
func (in *AWSControlPlane) DeepCopy() *AWSControlPlane {
	if in == nil {
		return nil
	}
	out := new(AWSControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSControlPlaneList) DeepCopyInto(out *AWSControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSControlPlaneList.
func (in *AWSControlPlaneList) DeepCopy() *AWSControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(AWSControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSControlPlaneSpec) DeepCopyInto(out *AWSControlPlaneSpec) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSControlPlaneSpec.
func (in *AWSControlPlaneSpec) DeepCopy() *AWSControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(AWSControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeployment) DeepCopyInto(out *AWSMachineDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
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
func (in *AWSMachineDeploymentSpecInstanceDistribution) DeepCopyInto(out *AWSMachineDeploymentSpecInstanceDistribution) {
	*out = *in
	if in.OnDemandPercentageAboveBaseCapacity != nil {
		in, out := &in.OnDemandPercentageAboveBaseCapacity, &out.OnDemandPercentageAboveBaseCapacity
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentSpecInstanceDistribution.
func (in *AWSMachineDeploymentSpecInstanceDistribution) DeepCopy() *AWSMachineDeploymentSpecInstanceDistribution {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentSpecInstanceDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentSpecNodePool) DeepCopyInto(out *AWSMachineDeploymentSpecNodePool) {
	*out = *in
	out.Machine = in.Machine
	out.Scaling = in.Scaling
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
	in.InstanceDistribution.DeepCopyInto(&out.InstanceDistribution)
	out.Worker = in.Worker
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
func (in *AWSMachineDeploymentStatus) DeepCopyInto(out *AWSMachineDeploymentStatus) {
	*out = *in
	in.Provider.DeepCopyInto(&out.Provider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentStatus.
func (in *AWSMachineDeploymentStatus) DeepCopy() *AWSMachineDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentStatusProvider) DeepCopyInto(out *AWSMachineDeploymentStatusProvider) {
	*out = *in
	in.Worker.DeepCopyInto(&out.Worker)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentStatusProvider.
func (in *AWSMachineDeploymentStatusProvider) DeepCopy() *AWSMachineDeploymentStatusProvider {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentStatusProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineDeploymentStatusProviderWorker) DeepCopyInto(out *AWSMachineDeploymentStatusProviderWorker) {
	*out = *in
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineDeploymentStatusProviderWorker.
func (in *AWSMachineDeploymentStatusProviderWorker) DeepCopy() *AWSMachineDeploymentStatusProviderWorker {
	if in == nil {
		return nil
	}
	out := new(AWSMachineDeploymentStatusProviderWorker)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlane) DeepCopyInto(out *G8sControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlane.
func (in *G8sControlPlane) DeepCopy() *G8sControlPlane {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *G8sControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlaneList) DeepCopyInto(out *G8sControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]G8sControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlaneList.
func (in *G8sControlPlaneList) DeepCopy() *G8sControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *G8sControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlaneSpec) DeepCopyInto(out *G8sControlPlaneSpec) {
	*out = *in
	out.InfrastructureRef = in.InfrastructureRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlaneSpec.
func (in *G8sControlPlaneSpec) DeepCopy() *G8sControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlaneStatus) DeepCopyInto(out *G8sControlPlaneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlaneStatus.
func (in *G8sControlPlaneStatus) DeepCopy() *G8sControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPool) DeepCopyInto(out *NetworkPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPool.
func (in *NetworkPool) DeepCopy() *NetworkPool {
	if in == nil {
		return nil
	}
	out := new(NetworkPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPoolList) DeepCopyInto(out *NetworkPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPoolList.
func (in *NetworkPoolList) DeepCopy() *NetworkPoolList {
	if in == nil {
		return nil
	}
	out := new(NetworkPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPoolSpec) DeepCopyInto(out *NetworkPoolSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPoolSpec.
func (in *NetworkPoolSpec) DeepCopy() *NetworkPoolSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPoolSpec)
	in.DeepCopyInto(out)
	return out
}
