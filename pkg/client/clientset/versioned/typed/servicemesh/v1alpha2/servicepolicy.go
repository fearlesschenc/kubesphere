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

	v1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/servicemesh/v1alpha2"
	scheme "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServicePoliciesGetter has a method to return a ServicePolicyInterface.
// A group's client should implement this interface.
type ServicePoliciesGetter interface {
	ServicePolicies(namespace string) ServicePolicyInterface
}

// ServicePolicyInterface has methods to work with ServicePolicy resources.
type ServicePolicyInterface interface {
	Create(*v1alpha2.ServicePolicy) (*v1alpha2.ServicePolicy, error)
	Update(*v1alpha2.ServicePolicy) (*v1alpha2.ServicePolicy, error)
	UpdateStatus(*v1alpha2.ServicePolicy) (*v1alpha2.ServicePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.ServicePolicy, error)
	List(opts v1.ListOptions) (*v1alpha2.ServicePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ServicePolicy, err error)
	ServicePolicyExpansion
}

// servicePolicies implements ServicePolicyInterface
type servicePolicies struct {
	client rest.Interface
	ns     string
}

// newServicePolicies returns a ServicePolicies
func newServicePolicies(c *ServicemeshV1alpha2Client, namespace string) *servicePolicies {
	return &servicePolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the servicePolicy, and returns the corresponding servicePolicy object, and an error if there is any.
func (c *servicePolicies) Get(name string, options v1.GetOptions) (result *v1alpha2.ServicePolicy, err error) {
	result = &v1alpha2.ServicePolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServicePolicies that match those selectors.
func (c *servicePolicies) List(opts v1.ListOptions) (result *v1alpha2.ServicePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ServicePolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested servicePolicies.
func (c *servicePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a servicePolicy and creates it.  Returns the server's representation of the servicePolicy, and an error, if there is any.
func (c *servicePolicies) Create(servicePolicy *v1alpha2.ServicePolicy) (result *v1alpha2.ServicePolicy, err error) {
	result = &v1alpha2.ServicePolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicepolicies").
		Body(servicePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a servicePolicy and updates it. Returns the server's representation of the servicePolicy, and an error, if there is any.
func (c *servicePolicies) Update(servicePolicy *v1alpha2.ServicePolicy) (result *v1alpha2.ServicePolicy, err error) {
	result = &v1alpha2.ServicePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicepolicies").
		Name(servicePolicy.Name).
		Body(servicePolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *servicePolicies) UpdateStatus(servicePolicy *v1alpha2.ServicePolicy) (result *v1alpha2.ServicePolicy, err error) {
	result = &v1alpha2.ServicePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicepolicies").
		Name(servicePolicy.Name).
		SubResource("status").
		Body(servicePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the servicePolicy and deletes it. Returns an error if one occurs.
func (c *servicePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *servicePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched servicePolicy.
func (c *servicePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ServicePolicy, err error) {
	result = &v1alpha2.ServicePolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
