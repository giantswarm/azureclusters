// +build !ignore_autogenerated

/*
Copyright 2020 Giant Swarm GmbH.

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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sBootstrapConfig) DeepCopyInto(out *G8sBootstrapConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sBootstrapConfig.
func (in *G8sBootstrapConfig) DeepCopy() *G8sBootstrapConfig {
	if in == nil {
		return nil
	}
	out := new(G8sBootstrapConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *G8sBootstrapConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sBootstrapConfigList) DeepCopyInto(out *G8sBootstrapConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]G8sBootstrapConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sBootstrapConfigList.
func (in *G8sBootstrapConfigList) DeepCopy() *G8sBootstrapConfigList {
	if in == nil {
		return nil
	}
	out := new(G8sBootstrapConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *G8sBootstrapConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sBootstrapSpec) DeepCopyInto(out *G8sBootstrapSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sBootstrapSpec.
func (in *G8sBootstrapSpec) DeepCopy() *G8sBootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(G8sBootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sBootstrapStatus) DeepCopyInto(out *G8sBootstrapStatus) {
	*out = *in
	if in.BootstrapData != nil {
		in, out := &in.BootstrapData, &out.BootstrapData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sBootstrapStatus.
func (in *G8sBootstrapStatus) DeepCopy() *G8sBootstrapStatus {
	if in == nil {
		return nil
	}
	out := new(G8sBootstrapStatus)
	in.DeepCopyInto(out)
	return out
}
