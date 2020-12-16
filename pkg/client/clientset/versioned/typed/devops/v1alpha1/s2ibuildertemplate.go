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

	v1alpha1 "github.com/fearlesschenc/kubesphere/pkg/apis/devops/v1alpha1"
	scheme "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// S2iBuilderTemplatesGetter has a method to return a S2iBuilderTemplateInterface.
// A group's client should implement this interface.
type S2iBuilderTemplatesGetter interface {
	S2iBuilderTemplates() S2iBuilderTemplateInterface
}

// S2iBuilderTemplateInterface has methods to work with S2iBuilderTemplate resources.
type S2iBuilderTemplateInterface interface {
	Create(*v1alpha1.S2iBuilderTemplate) (*v1alpha1.S2iBuilderTemplate, error)
	Update(*v1alpha1.S2iBuilderTemplate) (*v1alpha1.S2iBuilderTemplate, error)
	UpdateStatus(*v1alpha1.S2iBuilderTemplate) (*v1alpha1.S2iBuilderTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.S2iBuilderTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.S2iBuilderTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S2iBuilderTemplate, err error)
	S2iBuilderTemplateExpansion
}

// s2iBuilderTemplates implements S2iBuilderTemplateInterface
type s2iBuilderTemplates struct {
	client rest.Interface
}

// newS2iBuilderTemplates returns a S2iBuilderTemplates
func newS2iBuilderTemplates(c *DevopsV1alpha1Client) *s2iBuilderTemplates {
	return &s2iBuilderTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the s2iBuilderTemplate, and returns the corresponding s2iBuilderTemplate object, and an error if there is any.
func (c *s2iBuilderTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.S2iBuilderTemplate, err error) {
	result = &v1alpha1.S2iBuilderTemplate{}
	err = c.client.Get().
		Resource("s2ibuildertemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S2iBuilderTemplates that match those selectors.
func (c *s2iBuilderTemplates) List(opts v1.ListOptions) (result *v1alpha1.S2iBuilderTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S2iBuilderTemplateList{}
	err = c.client.Get().
		Resource("s2ibuildertemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s2iBuilderTemplates.
func (c *s2iBuilderTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("s2ibuildertemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a s2iBuilderTemplate and creates it.  Returns the server's representation of the s2iBuilderTemplate, and an error, if there is any.
func (c *s2iBuilderTemplates) Create(s2iBuilderTemplate *v1alpha1.S2iBuilderTemplate) (result *v1alpha1.S2iBuilderTemplate, err error) {
	result = &v1alpha1.S2iBuilderTemplate{}
	err = c.client.Post().
		Resource("s2ibuildertemplates").
		Body(s2iBuilderTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s2iBuilderTemplate and updates it. Returns the server's representation of the s2iBuilderTemplate, and an error, if there is any.
func (c *s2iBuilderTemplates) Update(s2iBuilderTemplate *v1alpha1.S2iBuilderTemplate) (result *v1alpha1.S2iBuilderTemplate, err error) {
	result = &v1alpha1.S2iBuilderTemplate{}
	err = c.client.Put().
		Resource("s2ibuildertemplates").
		Name(s2iBuilderTemplate.Name).
		Body(s2iBuilderTemplate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *s2iBuilderTemplates) UpdateStatus(s2iBuilderTemplate *v1alpha1.S2iBuilderTemplate) (result *v1alpha1.S2iBuilderTemplate, err error) {
	result = &v1alpha1.S2iBuilderTemplate{}
	err = c.client.Put().
		Resource("s2ibuildertemplates").
		Name(s2iBuilderTemplate.Name).
		SubResource("status").
		Body(s2iBuilderTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the s2iBuilderTemplate and deletes it. Returns an error if one occurs.
func (c *s2iBuilderTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("s2ibuildertemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s2iBuilderTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("s2ibuildertemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s2iBuilderTemplate.
func (c *s2iBuilderTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S2iBuilderTemplate, err error) {
	result = &v1alpha1.S2iBuilderTemplate{}
	err = c.client.Patch(pt).
		Resource("s2ibuildertemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
