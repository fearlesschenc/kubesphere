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

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "github.com/fearlesschenc/kubesphere/pkg/apis/types/v1beta1"
	scheme "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/scheme"
)

// FederatedClusterRoleBindingsGetter has a method to return a FederatedClusterRoleBindingInterface.
// A group's client should implement this interface.
type FederatedClusterRoleBindingsGetter interface {
	FederatedClusterRoleBindings(namespace string) FederatedClusterRoleBindingInterface
}

// FederatedClusterRoleBindingInterface has methods to work with FederatedClusterRoleBinding resources.
type FederatedClusterRoleBindingInterface interface {
	Create(*v1beta1.FederatedClusterRoleBinding) (*v1beta1.FederatedClusterRoleBinding, error)
	Update(*v1beta1.FederatedClusterRoleBinding) (*v1beta1.FederatedClusterRoleBinding, error)
	UpdateStatus(*v1beta1.FederatedClusterRoleBinding) (*v1beta1.FederatedClusterRoleBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedClusterRoleBinding, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedClusterRoleBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedClusterRoleBinding, err error)
	FederatedClusterRoleBindingExpansion
}

// federatedClusterRoleBindings implements FederatedClusterRoleBindingInterface
type federatedClusterRoleBindings struct {
	client rest.Interface
	ns     string
}

// newFederatedClusterRoleBindings returns a FederatedClusterRoleBindings
func newFederatedClusterRoleBindings(c *TypesV1beta1Client, namespace string) *federatedClusterRoleBindings {
	return &federatedClusterRoleBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedClusterRoleBinding, and returns the corresponding federatedClusterRoleBinding object, and an error if there is any.
func (c *federatedClusterRoleBindings) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	result = &v1beta1.FederatedClusterRoleBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedClusterRoleBindings that match those selectors.
func (c *federatedClusterRoleBindings) List(opts v1.ListOptions) (result *v1beta1.FederatedClusterRoleBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedClusterRoleBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedClusterRoleBindings.
func (c *federatedClusterRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a federatedClusterRoleBinding and creates it.  Returns the server's representation of the federatedClusterRoleBinding, and an error, if there is any.
func (c *federatedClusterRoleBindings) Create(federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	result = &v1beta1.FederatedClusterRoleBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		Body(federatedClusterRoleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedClusterRoleBinding and updates it. Returns the server's representation of the federatedClusterRoleBinding, and an error, if there is any.
func (c *federatedClusterRoleBindings) Update(federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	result = &v1beta1.FederatedClusterRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		Name(federatedClusterRoleBinding.Name).
		Body(federatedClusterRoleBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedClusterRoleBindings) UpdateStatus(federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	result = &v1beta1.FederatedClusterRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		Name(federatedClusterRoleBinding.Name).
		SubResource("status").
		Body(federatedClusterRoleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedClusterRoleBinding and deletes it. Returns an error if one occurs.
func (c *federatedClusterRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedClusterRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedClusterRoleBinding.
func (c *federatedClusterRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	result = &v1beta1.FederatedClusterRoleBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedclusterrolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
