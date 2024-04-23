//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 Giant Swarm GmbH.

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
	apiv1alpha3 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	cluster_apiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachinePool) DeepCopyInto(out *AzureMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachinePool.
func (in *AzureMachinePool) DeepCopy() *AzureMachinePool {
	if in == nil {
		return nil
	}
	out := new(AzureMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachinePoolInstanceStatus) DeepCopyInto(out *AzureMachinePoolInstanceStatus) {
	*out = *in
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(apiv1alpha3.VMState)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachinePoolInstanceStatus.
func (in *AzureMachinePoolInstanceStatus) DeepCopy() *AzureMachinePoolInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AzureMachinePoolInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachinePoolList) DeepCopyInto(out *AzureMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachinePoolList.
func (in *AzureMachinePoolList) DeepCopy() *AzureMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(AzureMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachinePoolSpec) DeepCopyInto(out *AzureMachinePoolSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]apiv1alpha3.UserAssignedIdentity, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachinePoolSpec.
func (in *AzureMachinePoolSpec) DeepCopy() *AzureMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(AzureMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachinePoolStatus) DeepCopyInto(out *AzureMachinePoolStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]*AzureMachinePoolInstanceStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AzureMachinePoolInstanceStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(apiv1alpha3.VMState)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(cluster_apiapiv1alpha3.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LongRunningOperationState != nil {
		in, out := &in.LongRunningOperationState, &out.LongRunningOperationState
		*out = new(apiv1alpha3.Future)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachinePoolStatus.
func (in *AzureMachinePoolStatus) DeepCopy() *AzureMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(AzureMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplate) DeepCopyInto(out *AzureMachineTemplate) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(apiv1alpha3.Image)
		(*in).DeepCopyInto(*out)
	}
	in.OSDisk.DeepCopyInto(&out.OSDisk)
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]apiv1alpha3.DataDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AcceleratedNetworking != nil {
		in, out := &in.AcceleratedNetworking, &out.AcceleratedNetworking
		*out = new(bool)
		**out = **in
	}
	if in.TerminateNotificationTimeout != nil {
		in, out := &in.TerminateNotificationTimeout, &out.TerminateNotificationTimeout
		*out = new(int)
		**out = **in
	}
	if in.SecurityProfile != nil {
		in, out := &in.SecurityProfile, &out.SecurityProfile
		*out = new(apiv1alpha3.SecurityProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.SpotVMOptions != nil {
		in, out := &in.SpotVMOptions, &out.SpotVMOptions
		*out = new(apiv1alpha3.SpotVMOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplate.
func (in *AzureMachineTemplate) DeepCopy() *AzureMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSS) DeepCopyInto(out *VMSS) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Image.DeepCopyInto(&out.Image)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(apiv1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]VMSSVM, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSS.
func (in *VMSS) DeepCopy() *VMSS {
	if in == nil {
		return nil
	}
	out := new(VMSS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSSVM) DeepCopyInto(out *VMSSVM) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSSVM.
func (in *VMSSVM) DeepCopy() *VMSSVM {
	if in == nil {
		return nil
	}
	out := new(VMSSVM)
	in.DeepCopyInto(out)
	return out
}
