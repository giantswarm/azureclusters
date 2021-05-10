// +build !ignore_autogenerated

/*
Copyright 2021 Giant Swarm GmbH.

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

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalog) DeepCopyInto(out *AppCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalog.
func (in *AppCatalog) DeepCopy() *AppCatalog {
	if in == nil {
		return nil
	}
	out := new(AppCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogList) DeepCopyInto(out *AppCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogList.
func (in *AppCatalogList) DeepCopy() *AppCatalogList {
	if in == nil {
		return nil
	}
	out := new(AppCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpec) DeepCopyInto(out *AppCatalogSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(AppCatalogSpecConfig)
		(*in).DeepCopyInto(*out)
	}
	out.Storage = in.Storage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpec.
func (in *AppCatalogSpec) DeepCopy() *AppCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecConfig) DeepCopyInto(out *AppCatalogSpecConfig) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(AppCatalogSpecConfigConfigMap)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(AppCatalogSpecConfigSecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecConfig.
func (in *AppCatalogSpecConfig) DeepCopy() *AppCatalogSpecConfig {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecConfigConfigMap) DeepCopyInto(out *AppCatalogSpecConfigConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecConfigConfigMap.
func (in *AppCatalogSpecConfigConfigMap) DeepCopy() *AppCatalogSpecConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecConfigSecret) DeepCopyInto(out *AppCatalogSpecConfigSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecConfigSecret.
func (in *AppCatalogSpecConfigSecret) DeepCopy() *AppCatalogSpecConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecStorage) DeepCopyInto(out *AppCatalogSpecStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecStorage.
func (in *AppCatalogSpecStorage) DeepCopy() *AppCatalogSpecStorage {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecStorage)
	in.DeepCopyInto(out)
	return out
}
