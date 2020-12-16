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

// WorkspaceRolesGetter has a method to return a WorkspaceRoleInterface.
// A group's client should implement this interface.
type WorkspaceRolesGetter interface {
	WorkspaceRoles() WorkspaceRoleInterface
}

// WorkspaceRoleInterface has methods to work with WorkspaceRole resources.
type WorkspaceRoleInterface interface {
	Create(*v1alpha2.WorkspaceRole) (*v1alpha2.WorkspaceRole, error)
	Update(*v1alpha2.WorkspaceRole) (*v1alpha2.WorkspaceRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.WorkspaceRole, error)
	List(opts v1.ListOptions) (*v1alpha2.WorkspaceRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.WorkspaceRole, err error)
	WorkspaceRoleExpansion
}

// workspaceRoles implements WorkspaceRoleInterface
type workspaceRoles struct {
	client rest.Interface
}

// newWorkspaceRoles returns a WorkspaceRoles
func newWorkspaceRoles(c *IamV1alpha2Client) *workspaceRoles {
	return &workspaceRoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the workspaceRole, and returns the corresponding workspaceRole object, and an error if there is any.
func (c *workspaceRoles) Get(name string, options v1.GetOptions) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Get().
		Resource("workspaceroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkspaceRoles that match those selectors.
func (c *workspaceRoles) List(opts v1.ListOptions) (result *v1alpha2.WorkspaceRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.WorkspaceRoleList{}
	err = c.client.Get().
		Resource("workspaceroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workspaceRoles.
func (c *workspaceRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("workspaceroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a workspaceRole and creates it.  Returns the server's representation of the workspaceRole, and an error, if there is any.
func (c *workspaceRoles) Create(workspaceRole *v1alpha2.WorkspaceRole) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Post().
		Resource("workspaceroles").
		Body(workspaceRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a workspaceRole and updates it. Returns the server's representation of the workspaceRole, and an error, if there is any.
func (c *workspaceRoles) Update(workspaceRole *v1alpha2.WorkspaceRole) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Put().
		Resource("workspaceroles").
		Name(workspaceRole.Name).
		Body(workspaceRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the workspaceRole and deletes it. Returns an error if one occurs.
func (c *workspaceRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("workspaceroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workspaceRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("workspaceroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched workspaceRole.
func (c *workspaceRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Patch(pt).
		Resource("workspaceroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
