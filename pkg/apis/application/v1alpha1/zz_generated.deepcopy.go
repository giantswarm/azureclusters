//go:build !ignore_autogenerated

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *App) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalog) DeepCopyInto(out *AppCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
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
func (in *AppCatalogEntry) DeepCopyInto(out *AppCatalogEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogEntry.
func (in *AppCatalogEntry) DeepCopy() *AppCatalogEntry {
	if in == nil {
		return nil
	}
	out := new(AppCatalogEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppCatalogEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogEntryList) DeepCopyInto(out *AppCatalogEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppCatalogEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogEntryList.
func (in *AppCatalogEntryList) DeepCopy() *AppCatalogEntryList {
	if in == nil {
		return nil
	}
	out := new(AppCatalogEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppCatalogEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogEntrySpec) DeepCopyInto(out *AppCatalogEntrySpec) {
	*out = *in
	out.Catalog = in.Catalog
	in.Chart.DeepCopyInto(&out.Chart)
	if in.DateCreated != nil {
		in, out := &in.DateCreated, &out.DateCreated
		*out = (*in).DeepCopy()
	}
	if in.DateUpdated != nil {
		in, out := &in.DateUpdated, &out.DateUpdated
		*out = (*in).DeepCopy()
	}
	if in.Restrictions != nil {
		in, out := &in.Restrictions, &out.Restrictions
		*out = new(AppCatalogEntrySpecRestrictions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogEntrySpec.
func (in *AppCatalogEntrySpec) DeepCopy() *AppCatalogEntrySpec {
	if in == nil {
		return nil
	}
	out := new(AppCatalogEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogEntrySpecCatalog) DeepCopyInto(out *AppCatalogEntrySpecCatalog) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogEntrySpecCatalog.
func (in *AppCatalogEntrySpecCatalog) DeepCopy() *AppCatalogEntrySpecCatalog {
	if in == nil {
		return nil
	}
	out := new(AppCatalogEntrySpecCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogEntrySpecChart) DeepCopyInto(out *AppCatalogEntrySpecChart) {
	*out = *in
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogEntrySpecChart.
func (in *AppCatalogEntrySpecChart) DeepCopy() *AppCatalogEntrySpecChart {
	if in == nil {
		return nil
	}
	out := new(AppCatalogEntrySpecChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogEntrySpecRestrictions) DeepCopyInto(out *AppCatalogEntrySpecRestrictions) {
	*out = *in
	if in.CompatibleProviders != nil {
		in, out := &in.CompatibleProviders, &out.CompatibleProviders
		*out = make([]Provider, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogEntrySpecRestrictions.
func (in *AppCatalogEntrySpecRestrictions) DeepCopy() *AppCatalogEntrySpecRestrictions {
	if in == nil {
		return nil
	}
	out := new(AppCatalogEntrySpecRestrictions)
	in.DeepCopyInto(out)
	return out
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
	out.Config = in.Config
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
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppList) DeepCopyInto(out *AppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]App, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppList.
func (in *AppList) DeepCopy() *AppList {
	if in == nil {
		return nil
	}
	out := new(AppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	out.Config = in.Config
	out.Install = in.Install
	out.KubeConfig = in.KubeConfig
	in.NamespaceConfig.DeepCopyInto(&out.NamespaceConfig)
	out.UserConfig = in.UserConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec.
func (in *AppSpec) DeepCopy() *AppSpec {
	if in == nil {
		return nil
	}
	out := new(AppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecConfig) DeepCopyInto(out *AppSpecConfig) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecConfig.
func (in *AppSpecConfig) DeepCopy() *AppSpecConfig {
	if in == nil {
		return nil
	}
	out := new(AppSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecConfigConfigMap) DeepCopyInto(out *AppSpecConfigConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecConfigConfigMap.
func (in *AppSpecConfigConfigMap) DeepCopy() *AppSpecConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(AppSpecConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecConfigSecret) DeepCopyInto(out *AppSpecConfigSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecConfigSecret.
func (in *AppSpecConfigSecret) DeepCopy() *AppSpecConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppSpecConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecInstall) DeepCopyInto(out *AppSpecInstall) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecInstall.
func (in *AppSpecInstall) DeepCopy() *AppSpecInstall {
	if in == nil {
		return nil
	}
	out := new(AppSpecInstall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecKubeConfig) DeepCopyInto(out *AppSpecKubeConfig) {
	*out = *in
	out.Context = in.Context
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecKubeConfig.
func (in *AppSpecKubeConfig) DeepCopy() *AppSpecKubeConfig {
	if in == nil {
		return nil
	}
	out := new(AppSpecKubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecKubeConfigContext) DeepCopyInto(out *AppSpecKubeConfigContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecKubeConfigContext.
func (in *AppSpecKubeConfigContext) DeepCopy() *AppSpecKubeConfigContext {
	if in == nil {
		return nil
	}
	out := new(AppSpecKubeConfigContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecKubeConfigSecret) DeepCopyInto(out *AppSpecKubeConfigSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecKubeConfigSecret.
func (in *AppSpecKubeConfigSecret) DeepCopy() *AppSpecKubeConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppSpecKubeConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecNamespaceConfig) DeepCopyInto(out *AppSpecNamespaceConfig) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecNamespaceConfig.
func (in *AppSpecNamespaceConfig) DeepCopy() *AppSpecNamespaceConfig {
	if in == nil {
		return nil
	}
	out := new(AppSpecNamespaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecUserConfig) DeepCopyInto(out *AppSpecUserConfig) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecUserConfig.
func (in *AppSpecUserConfig) DeepCopy() *AppSpecUserConfig {
	if in == nil {
		return nil
	}
	out := new(AppSpecUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecUserConfigConfigMap) DeepCopyInto(out *AppSpecUserConfigConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecUserConfigConfigMap.
func (in *AppSpecUserConfigConfigMap) DeepCopy() *AppSpecUserConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(AppSpecUserConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecUserConfigSecret) DeepCopyInto(out *AppSpecUserConfigSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecUserConfigSecret.
func (in *AppSpecUserConfigSecret) DeepCopy() *AppSpecUserConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppSpecUserConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	in.Release.DeepCopyInto(&out.Release)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatusRelease) DeepCopyInto(out *AppStatusRelease) {
	*out = *in
	in.LastDeployed.DeepCopyInto(&out.LastDeployed)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatusRelease.
func (in *AppStatusRelease) DeepCopy() *AppStatusRelease {
	if in == nil {
		return nil
	}
	out := new(AppStatusRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Catalog) DeepCopyInto(out *Catalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Catalog.
func (in *Catalog) DeepCopy() *Catalog {
	if in == nil {
		return nil
	}
	out := new(Catalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Catalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogList) DeepCopyInto(out *CatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Catalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogList.
func (in *CatalogList) DeepCopy() *CatalogList {
	if in == nil {
		return nil
	}
	out := new(CatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpec) DeepCopyInto(out *CatalogSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(CatalogSpecConfig)
		(*in).DeepCopyInto(*out)
	}
	out.Storage = in.Storage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpec.
func (in *CatalogSpec) DeepCopy() *CatalogSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpecConfig) DeepCopyInto(out *CatalogSpecConfig) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(CatalogSpecConfigConfigMap)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(CatalogSpecConfigSecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpecConfig.
func (in *CatalogSpecConfig) DeepCopy() *CatalogSpecConfig {
	if in == nil {
		return nil
	}
	out := new(CatalogSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpecConfigConfigMap) DeepCopyInto(out *CatalogSpecConfigConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpecConfigConfigMap.
func (in *CatalogSpecConfigConfigMap) DeepCopy() *CatalogSpecConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(CatalogSpecConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpecConfigSecret) DeepCopyInto(out *CatalogSpecConfigSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpecConfigSecret.
func (in *CatalogSpecConfigSecret) DeepCopy() *CatalogSpecConfigSecret {
	if in == nil {
		return nil
	}
	out := new(CatalogSpecConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpecStorage) DeepCopyInto(out *CatalogSpecStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpecStorage.
func (in *CatalogSpecStorage) DeepCopy() *CatalogSpecStorage {
	if in == nil {
		return nil
	}
	out := new(CatalogSpecStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Chart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartList) DeepCopyInto(out *ChartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Chart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartList.
func (in *ChartList) DeepCopy() *ChartList {
	if in == nil {
		return nil
	}
	out := new(ChartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpec) DeepCopyInto(out *ChartSpec) {
	*out = *in
	out.Config = in.Config
	out.Install = in.Install
	in.NamespaceConfig.DeepCopyInto(&out.NamespaceConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpec.
func (in *ChartSpec) DeepCopy() *ChartSpec {
	if in == nil {
		return nil
	}
	out := new(ChartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecConfig) DeepCopyInto(out *ChartSpecConfig) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecConfig.
func (in *ChartSpecConfig) DeepCopy() *ChartSpecConfig {
	if in == nil {
		return nil
	}
	out := new(ChartSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecConfigConfigMap) DeepCopyInto(out *ChartSpecConfigConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecConfigConfigMap.
func (in *ChartSpecConfigConfigMap) DeepCopy() *ChartSpecConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(ChartSpecConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecConfigSecret) DeepCopyInto(out *ChartSpecConfigSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecConfigSecret.
func (in *ChartSpecConfigSecret) DeepCopy() *ChartSpecConfigSecret {
	if in == nil {
		return nil
	}
	out := new(ChartSpecConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecInstall) DeepCopyInto(out *ChartSpecInstall) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecInstall.
func (in *ChartSpecInstall) DeepCopy() *ChartSpecInstall {
	if in == nil {
		return nil
	}
	out := new(ChartSpecInstall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecNamespaceConfig) DeepCopyInto(out *ChartSpecNamespaceConfig) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecNamespaceConfig.
func (in *ChartSpecNamespaceConfig) DeepCopy() *ChartSpecNamespaceConfig {
	if in == nil {
		return nil
	}
	out := new(ChartSpecNamespaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartStatus) DeepCopyInto(out *ChartStatus) {
	*out = *in
	in.Release.DeepCopyInto(&out.Release)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartStatus.
func (in *ChartStatus) DeepCopy() *ChartStatus {
	if in == nil {
		return nil
	}
	out := new(ChartStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartStatusRelease) DeepCopyInto(out *ChartStatusRelease) {
	*out = *in
	if in.LastDeployed != nil {
		in, out := &in.LastDeployed, &out.LastDeployed
		*out = (*in).DeepCopy()
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartStatusRelease.
func (in *ChartStatusRelease) DeepCopy() *ChartStatusRelease {
	if in == nil {
		return nil
	}
	out := new(ChartStatusRelease)
	in.DeepCopyInto(out)
	return out
}
