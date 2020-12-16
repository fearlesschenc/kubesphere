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

// FederatedPersistentVolumeClaimsGetter has a method to return a FederatedPersistentVolumeClaimInterface.
// A group's client should implement this interface.
type FederatedPersistentVolumeClaimsGetter interface {
	FederatedPersistentVolumeClaims(namespace string) FederatedPersistentVolumeClaimInterface
}

// FederatedPersistentVolumeClaimInterface has methods to work with FederatedPersistentVolumeClaim resources.
type FederatedPersistentVolumeClaimInterface interface {
	Create(*v1beta1.FederatedPersistentVolumeClaim) (*v1beta1.FederatedPersistentVolumeClaim, error)
	Update(*v1beta1.FederatedPersistentVolumeClaim) (*v1beta1.FederatedPersistentVolumeClaim, error)
	UpdateStatus(*v1beta1.FederatedPersistentVolumeClaim) (*v1beta1.FederatedPersistentVolumeClaim, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedPersistentVolumeClaim, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedPersistentVolumeClaimList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedPersistentVolumeClaim, err error)
	FederatedPersistentVolumeClaimExpansion
}

// federatedPersistentVolumeClaims implements FederatedPersistentVolumeClaimInterface
type federatedPersistentVolumeClaims struct {
	client rest.Interface
	ns     string
}

// newFederatedPersistentVolumeClaims returns a FederatedPersistentVolumeClaims
func newFederatedPersistentVolumeClaims(c *TypesV1beta1Client, namespace string) *federatedPersistentVolumeClaims {
	return &federatedPersistentVolumeClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedPersistentVolumeClaim, and returns the corresponding federatedPersistentVolumeClaim object, and an error if there is any.
func (c *federatedPersistentVolumeClaims) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedPersistentVolumeClaims that match those selectors.
func (c *federatedPersistentVolumeClaims) List(opts v1.ListOptions) (result *v1beta1.FederatedPersistentVolumeClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedPersistentVolumeClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedPersistentVolumeClaims.
func (c *federatedPersistentVolumeClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a federatedPersistentVolumeClaim and creates it.  Returns the server's representation of the federatedPersistentVolumeClaim, and an error, if there is any.
func (c *federatedPersistentVolumeClaims) Create(federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Body(federatedPersistentVolumeClaim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedPersistentVolumeClaim and updates it. Returns the server's representation of the federatedPersistentVolumeClaim, and an error, if there is any.
func (c *federatedPersistentVolumeClaims) Update(federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(federatedPersistentVolumeClaim.Name).
		Body(federatedPersistentVolumeClaim).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedPersistentVolumeClaims) UpdateStatus(federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(federatedPersistentVolumeClaim.Name).
		SubResource("status").
		Body(federatedPersistentVolumeClaim).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedPersistentVolumeClaim and deletes it. Returns an error if one occurs.
func (c *federatedPersistentVolumeClaims) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedPersistentVolumeClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedPersistentVolumeClaim.
func (c *federatedPersistentVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
