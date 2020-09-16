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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
)

// FakeAWSIPPools implements AWSIPPoolInterface
type FakeAWSIPPools struct {
	Fake *FakeInfrastructureV1alpha2
	ns   string
}

var awsippoolsResource = schema.GroupVersionResource{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Resource: "awsippools"}

var awsippoolsKind = schema.GroupVersionKind{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Kind: "AWSIPPool"}

// Get takes name of the aWSIPPool, and returns the corresponding aWSIPPool object, and an error if there is any.
func (c *FakeAWSIPPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.AWSIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsippoolsResource, c.ns, name), &v1alpha2.AWSIPPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSIPPool), err
}

// List takes label and field selectors, and returns the list of AWSIPPools that match those selectors.
func (c *FakeAWSIPPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.AWSIPPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsippoolsResource, awsippoolsKind, c.ns, opts), &v1alpha2.AWSIPPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.AWSIPPoolList{ListMeta: obj.(*v1alpha2.AWSIPPoolList).ListMeta}
	for _, item := range obj.(*v1alpha2.AWSIPPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSIPPools.
func (c *FakeAWSIPPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsippoolsResource, c.ns, opts))

}

// Create takes the representation of a aWSIPPool and creates it.  Returns the server's representation of the aWSIPPool, and an error, if there is any.
func (c *FakeAWSIPPools) Create(ctx context.Context, aWSIPPool *v1alpha2.AWSIPPool, opts v1.CreateOptions) (result *v1alpha2.AWSIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsippoolsResource, c.ns, aWSIPPool), &v1alpha2.AWSIPPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSIPPool), err
}

// Update takes the representation of a aWSIPPool and updates it. Returns the server's representation of the aWSIPPool, and an error, if there is any.
func (c *FakeAWSIPPools) Update(ctx context.Context, aWSIPPool *v1alpha2.AWSIPPool, opts v1.UpdateOptions) (result *v1alpha2.AWSIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsippoolsResource, c.ns, aWSIPPool), &v1alpha2.AWSIPPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSIPPool), err
}

// Delete takes name of the aWSIPPool and deletes it. Returns an error if one occurs.
func (c *FakeAWSIPPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsippoolsResource, c.ns, name), &v1alpha2.AWSIPPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSIPPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsippoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.AWSIPPoolList{})
	return err
}

// Patch applies the patch and returns the patched aWSIPPool.
func (c *FakeAWSIPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.AWSIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsippoolsResource, c.ns, name, pt, data, subresources...), &v1alpha2.AWSIPPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSIPPool), err
}
