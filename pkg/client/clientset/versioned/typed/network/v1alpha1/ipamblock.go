/*
Copyright 2020 The KubeSphere Authors.

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
	"time"

	v1alpha1 "github.com/fearlesschenc/kubesphere/pkg/apis/network/v1alpha1"
	scheme "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IPAMBlocksGetter has a method to return a IPAMBlockInterface.
// A group's client should implement this interface.
type IPAMBlocksGetter interface {
	IPAMBlocks() IPAMBlockInterface
}

// IPAMBlockInterface has methods to work with IPAMBlock resources.
type IPAMBlockInterface interface {
	Create(*v1alpha1.IPAMBlock) (*v1alpha1.IPAMBlock, error)
	Update(*v1alpha1.IPAMBlock) (*v1alpha1.IPAMBlock, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IPAMBlock, error)
	List(opts v1.ListOptions) (*v1alpha1.IPAMBlockList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IPAMBlock, err error)
	IPAMBlockExpansion
}

// iPAMBlocks implements IPAMBlockInterface
type iPAMBlocks struct {
	client rest.Interface
}

// newIPAMBlocks returns a IPAMBlocks
func newIPAMBlocks(c *NetworkV1alpha1Client) *iPAMBlocks {
	return &iPAMBlocks{
		client: c.RESTClient(),
	}
}

// Get takes name of the iPAMBlock, and returns the corresponding iPAMBlock object, and an error if there is any.
func (c *iPAMBlocks) Get(name string, options v1.GetOptions) (result *v1alpha1.IPAMBlock, err error) {
	result = &v1alpha1.IPAMBlock{}
	err = c.client.Get().
		Resource("ipamblocks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IPAMBlocks that match those selectors.
func (c *iPAMBlocks) List(opts v1.ListOptions) (result *v1alpha1.IPAMBlockList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IPAMBlockList{}
	err = c.client.Get().
		Resource("ipamblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iPAMBlocks.
func (c *iPAMBlocks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ipamblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iPAMBlock and creates it.  Returns the server's representation of the iPAMBlock, and an error, if there is any.
func (c *iPAMBlocks) Create(iPAMBlock *v1alpha1.IPAMBlock) (result *v1alpha1.IPAMBlock, err error) {
	result = &v1alpha1.IPAMBlock{}
	err = c.client.Post().
		Resource("ipamblocks").
		Body(iPAMBlock).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iPAMBlock and updates it. Returns the server's representation of the iPAMBlock, and an error, if there is any.
func (c *iPAMBlocks) Update(iPAMBlock *v1alpha1.IPAMBlock) (result *v1alpha1.IPAMBlock, err error) {
	result = &v1alpha1.IPAMBlock{}
	err = c.client.Put().
		Resource("ipamblocks").
		Name(iPAMBlock.Name).
		Body(iPAMBlock).
		Do().
		Into(result)
	return
}

// Delete takes name of the iPAMBlock and deletes it. Returns an error if one occurs.
func (c *iPAMBlocks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ipamblocks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iPAMBlocks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ipamblocks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iPAMBlock.
func (c *iPAMBlocks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IPAMBlock, err error) {
	result = &v1alpha1.IPAMBlock{}
	err = c.client.Patch(pt).
		Resource("ipamblocks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
