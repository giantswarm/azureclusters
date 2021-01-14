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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKVMConfigs implements KVMConfigInterface
type FakeKVMConfigs struct {
	Fake *FakeProviderV1alpha1
	ns   string
}

var kvmconfigsResource = schema.GroupVersionResource{Group: "provider.giantswarm.io", Version: "v1alpha1", Resource: "kvmconfigs"}

var kvmconfigsKind = schema.GroupVersionKind{Group: "provider.giantswarm.io", Version: "v1alpha1", Kind: "KVMConfig"}

// Get takes name of the kVMConfig, and returns the corresponding kVMConfig object, and an error if there is any.
func (c *FakeKVMConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KVMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kvmconfigsResource, c.ns, name), &v1alpha1.KVMConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KVMConfig), err
}

// List takes label and field selectors, and returns the list of KVMConfigs that match those selectors.
func (c *FakeKVMConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KVMConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kvmconfigsResource, kvmconfigsKind, c.ns, opts), &v1alpha1.KVMConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KVMConfigList{ListMeta: obj.(*v1alpha1.KVMConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.KVMConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kVMConfigs.
func (c *FakeKVMConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kvmconfigsResource, c.ns, opts))

}

// Create takes the representation of a kVMConfig and creates it.  Returns the server's representation of the kVMConfig, and an error, if there is any.
func (c *FakeKVMConfigs) Create(ctx context.Context, kVMConfig *v1alpha1.KVMConfig, opts v1.CreateOptions) (result *v1alpha1.KVMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kvmconfigsResource, c.ns, kVMConfig), &v1alpha1.KVMConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KVMConfig), err
}

// Update takes the representation of a kVMConfig and updates it. Returns the server's representation of the kVMConfig, and an error, if there is any.
func (c *FakeKVMConfigs) Update(ctx context.Context, kVMConfig *v1alpha1.KVMConfig, opts v1.UpdateOptions) (result *v1alpha1.KVMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kvmconfigsResource, c.ns, kVMConfig), &v1alpha1.KVMConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KVMConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKVMConfigs) UpdateStatus(ctx context.Context, kVMConfig *v1alpha1.KVMConfig, opts v1.UpdateOptions) (*v1alpha1.KVMConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kvmconfigsResource, "status", c.ns, kVMConfig), &v1alpha1.KVMConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KVMConfig), err
}

// Delete takes name of the kVMConfig and deletes it. Returns an error if one occurs.
func (c *FakeKVMConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kvmconfigsResource, c.ns, name), &v1alpha1.KVMConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKVMConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kvmconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KVMConfigList{})
	return err
}

// Patch applies the patch and returns the patched kVMConfig.
func (c *FakeKVMConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KVMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kvmconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KVMConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KVMConfig), err
}
