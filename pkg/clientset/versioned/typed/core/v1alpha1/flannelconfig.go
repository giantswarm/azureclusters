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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/core/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/v3/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FlannelConfigsGetter has a method to return a FlannelConfigInterface.
// A group's client should implement this interface.
type FlannelConfigsGetter interface {
	FlannelConfigs(namespace string) FlannelConfigInterface
}

// FlannelConfigInterface has methods to work with FlannelConfig resources.
type FlannelConfigInterface interface {
	Create(ctx context.Context, flannelConfig *v1alpha1.FlannelConfig, opts v1.CreateOptions) (*v1alpha1.FlannelConfig, error)
	Update(ctx context.Context, flannelConfig *v1alpha1.FlannelConfig, opts v1.UpdateOptions) (*v1alpha1.FlannelConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FlannelConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FlannelConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FlannelConfig, err error)
	FlannelConfigExpansion
}

// flannelConfigs implements FlannelConfigInterface
type flannelConfigs struct {
	client rest.Interface
	ns     string
}

// newFlannelConfigs returns a FlannelConfigs
func newFlannelConfigs(c *CoreV1alpha1Client, namespace string) *flannelConfigs {
	return &flannelConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the flannelConfig, and returns the corresponding flannelConfig object, and an error if there is any.
func (c *flannelConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FlannelConfig, err error) {
	result = &v1alpha1.FlannelConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flannelconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FlannelConfigs that match those selectors.
func (c *flannelConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FlannelConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FlannelConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flannelconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested flannelConfigs.
func (c *flannelConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("flannelconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a flannelConfig and creates it.  Returns the server's representation of the flannelConfig, and an error, if there is any.
func (c *flannelConfigs) Create(ctx context.Context, flannelConfig *v1alpha1.FlannelConfig, opts v1.CreateOptions) (result *v1alpha1.FlannelConfig, err error) {
	result = &v1alpha1.FlannelConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("flannelconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flannelConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a flannelConfig and updates it. Returns the server's representation of the flannelConfig, and an error, if there is any.
func (c *flannelConfigs) Update(ctx context.Context, flannelConfig *v1alpha1.FlannelConfig, opts v1.UpdateOptions) (result *v1alpha1.FlannelConfig, err error) {
	result = &v1alpha1.FlannelConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flannelconfigs").
		Name(flannelConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flannelConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the flannelConfig and deletes it. Returns an error if one occurs.
func (c *flannelConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flannelconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flannelConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flannelconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched flannelConfig.
func (c *flannelConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FlannelConfig, err error) {
	result = &v1alpha1.FlannelConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("flannelconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
