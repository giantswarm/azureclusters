/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeConfigs implements NodeConfigInterface
type FakeNodeConfigs struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var nodeconfigsResource = schema.GroupVersionResource{Group: "core.giantswarm.io", Version: "v1alpha1", Resource: "nodeconfigs"}

var nodeconfigsKind = schema.GroupVersionKind{Group: "core.giantswarm.io", Version: "v1alpha1", Kind: "NodeConfig"}

// Get takes name of the nodeConfig, and returns the corresponding nodeConfig object, and an error if there is any.
func (c *FakeNodeConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodeconfigsResource, c.ns, name), &v1alpha1.NodeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}

// List takes label and field selectors, and returns the list of NodeConfigs that match those selectors.
func (c *FakeNodeConfigs) List(opts v1.ListOptions) (result *v1alpha1.NodeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodeconfigsResource, nodeconfigsKind, c.ns, opts), &v1alpha1.NodeConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeConfigList{}
	for _, item := range obj.(*v1alpha1.NodeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeConfigs.
func (c *FakeNodeConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodeconfigsResource, c.ns, opts))

}

// Create takes the representation of a nodeConfig and creates it.  Returns the server's representation of the nodeConfig, and an error, if there is any.
func (c *FakeNodeConfigs) Create(nodeConfig *v1alpha1.NodeConfig) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodeconfigsResource, c.ns, nodeConfig), &v1alpha1.NodeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}

// Update takes the representation of a nodeConfig and updates it. Returns the server's representation of the nodeConfig, and an error, if there is any.
func (c *FakeNodeConfigs) Update(nodeConfig *v1alpha1.NodeConfig) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodeconfigsResource, c.ns, nodeConfig), &v1alpha1.NodeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}

// Delete takes name of the nodeConfig and deletes it. Returns an error if one occurs.
func (c *FakeNodeConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodeconfigsResource, c.ns, name), &v1alpha1.NodeConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodeconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeConfigList{})
	return err
}

// Patch applies the patch and returns the patched nodeConfig.
func (c *FakeNodeConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodeconfigsResource, c.ns, name, data, subresources...), &v1alpha1.NodeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}
