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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "github.com/fearlesschenc/kubesphere/pkg/apis/tenant/v1alpha1"
	scheme "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/scheme"
)

// WorkspacesGetter has a method to return a WorkspaceInterface.
// A group's client should implement this interface.
type WorkspacesGetter interface {
	Workspaces() WorkspaceInterface
}

// WorkspaceInterface has methods to work with Workspace resources.
type WorkspaceInterface interface {
	Create(*v1alpha1.Workspace) (*v1alpha1.Workspace, error)
	Update(*v1alpha1.Workspace) (*v1alpha1.Workspace, error)
	UpdateStatus(*v1alpha1.Workspace) (*v1alpha1.Workspace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Workspace, error)
	List(opts v1.ListOptions) (*v1alpha1.WorkspaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Workspace, err error)
	WorkspaceExpansion
}

// workspaces implements WorkspaceInterface
type workspaces struct {
	client rest.Interface
}

// newWorkspaces returns a Workspaces
func newWorkspaces(c *TenantV1alpha1Client) *workspaces {
	return &workspaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the workspace, and returns the corresponding workspace object, and an error if there is any.
func (c *workspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.Workspace, err error) {
	result = &v1alpha1.Workspace{}
	err = c.client.Get().
		Resource("workspaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Workspaces that match those selectors.
func (c *workspaces) List(opts v1.ListOptions) (result *v1alpha1.WorkspaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WorkspaceList{}
	err = c.client.Get().
		Resource("workspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workspaces.
func (c *workspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("workspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a workspace and creates it.  Returns the server's representation of the workspace, and an error, if there is any.
func (c *workspaces) Create(workspace *v1alpha1.Workspace) (result *v1alpha1.Workspace, err error) {
	result = &v1alpha1.Workspace{}
	err = c.client.Post().
		Resource("workspaces").
		Body(workspace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a workspace and updates it. Returns the server's representation of the workspace, and an error, if there is any.
func (c *workspaces) Update(workspace *v1alpha1.Workspace) (result *v1alpha1.Workspace, err error) {
	result = &v1alpha1.Workspace{}
	err = c.client.Put().
		Resource("workspaces").
		Name(workspace.Name).
		Body(workspace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *workspaces) UpdateStatus(workspace *v1alpha1.Workspace) (result *v1alpha1.Workspace, err error) {
	result = &v1alpha1.Workspace{}
	err = c.client.Put().
		Resource("workspaces").
		Name(workspace.Name).
		SubResource("status").
		Body(workspace).
		Do().
		Into(result)
	return
}

// Delete takes name of the workspace and deletes it. Returns an error if one occurs.
func (c *workspaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("workspaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("workspaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched workspace.
func (c *workspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Workspace, err error) {
	result = &v1alpha1.Workspace{}
	err = c.client.Patch(pt).
		Resource("workspaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
