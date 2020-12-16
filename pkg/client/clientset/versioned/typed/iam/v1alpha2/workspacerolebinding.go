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

package v1alpha2

import (
	"time"

	v1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/iam/v1alpha2"
	scheme "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkspaceRoleBindingsGetter has a method to return a WorkspaceRoleBindingInterface.
// A group's client should implement this interface.
type WorkspaceRoleBindingsGetter interface {
	WorkspaceRoleBindings() WorkspaceRoleBindingInterface
}

// WorkspaceRoleBindingInterface has methods to work with WorkspaceRoleBinding resources.
type WorkspaceRoleBindingInterface interface {
	Create(*v1alpha2.WorkspaceRoleBinding) (*v1alpha2.WorkspaceRoleBinding, error)
	Update(*v1alpha2.WorkspaceRoleBinding) (*v1alpha2.WorkspaceRoleBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.WorkspaceRoleBinding, error)
	List(opts v1.ListOptions) (*v1alpha2.WorkspaceRoleBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.WorkspaceRoleBinding, err error)
	WorkspaceRoleBindingExpansion
}

// workspaceRoleBindings implements WorkspaceRoleBindingInterface
type workspaceRoleBindings struct {
	client rest.Interface
}

// newWorkspaceRoleBindings returns a WorkspaceRoleBindings
func newWorkspaceRoleBindings(c *IamV1alpha2Client) *workspaceRoleBindings {
	return &workspaceRoleBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the workspaceRoleBinding, and returns the corresponding workspaceRoleBinding object, and an error if there is any.
func (c *workspaceRoleBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.WorkspaceRoleBinding, err error) {
	result = &v1alpha2.WorkspaceRoleBinding{}
	err = c.client.Get().
		Resource("workspacerolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkspaceRoleBindings that match those selectors.
func (c *workspaceRoleBindings) List(opts v1.ListOptions) (result *v1alpha2.WorkspaceRoleBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.WorkspaceRoleBindingList{}
	err = c.client.Get().
		Resource("workspacerolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workspaceRoleBindings.
func (c *workspaceRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("workspacerolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a workspaceRoleBinding and creates it.  Returns the server's representation of the workspaceRoleBinding, and an error, if there is any.
func (c *workspaceRoleBindings) Create(workspaceRoleBinding *v1alpha2.WorkspaceRoleBinding) (result *v1alpha2.WorkspaceRoleBinding, err error) {
	result = &v1alpha2.WorkspaceRoleBinding{}
	err = c.client.Post().
		Resource("workspacerolebindings").
		Body(workspaceRoleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a workspaceRoleBinding and updates it. Returns the server's representation of the workspaceRoleBinding, and an error, if there is any.
func (c *workspaceRoleBindings) Update(workspaceRoleBinding *v1alpha2.WorkspaceRoleBinding) (result *v1alpha2.WorkspaceRoleBinding, err error) {
	result = &v1alpha2.WorkspaceRoleBinding{}
	err = c.client.Put().
		Resource("workspacerolebindings").
		Name(workspaceRoleBinding.Name).
		Body(workspaceRoleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the workspaceRoleBinding and deletes it. Returns an error if one occurs.
func (c *workspaceRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("workspacerolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workspaceRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("workspacerolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched workspaceRoleBinding.
func (c *workspaceRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.WorkspaceRoleBinding, err error) {
	result = &v1alpha2.WorkspaceRoleBinding{}
	err = c.client.Patch(pt).
		Resource("workspacerolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
