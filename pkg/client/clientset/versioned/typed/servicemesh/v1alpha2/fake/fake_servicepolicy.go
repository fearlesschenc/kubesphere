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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/servicemesh/v1alpha2"
)

// FakeServicePolicies implements ServicePolicyInterface
type FakeServicePolicies struct {
	Fake *FakeServicemeshV1alpha2
	ns   string
}

var servicepoliciesResource = schema.GroupVersionResource{Group: "servicemesh.kubesphere.io", Version: "v1alpha2", Resource: "servicepolicies"}

var servicepoliciesKind = schema.GroupVersionKind{Group: "servicemesh.kubesphere.io", Version: "v1alpha2", Kind: "ServicePolicy"}

// Get takes name of the servicePolicy, and returns the corresponding servicePolicy object, and an error if there is any.
func (c *FakeServicePolicies) Get(name string, options v1.GetOptions) (result *v1alpha2.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicepoliciesResource, c.ns, name), &v1alpha2.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ServicePolicy), err
}

// List takes label and field selectors, and returns the list of ServicePolicies that match those selectors.
func (c *FakeServicePolicies) List(opts v1.ListOptions) (result *v1alpha2.ServicePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicepoliciesResource, servicepoliciesKind, c.ns, opts), &v1alpha2.ServicePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ServicePolicyList{ListMeta: obj.(*v1alpha2.ServicePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha2.ServicePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicePolicies.
func (c *FakeServicePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicepoliciesResource, c.ns, opts))

}

// Create takes the representation of a servicePolicy and creates it.  Returns the server's representation of the servicePolicy, and an error, if there is any.
func (c *FakeServicePolicies) Create(servicePolicy *v1alpha2.ServicePolicy) (result *v1alpha2.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicepoliciesResource, c.ns, servicePolicy), &v1alpha2.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ServicePolicy), err
}

// Update takes the representation of a servicePolicy and updates it. Returns the server's representation of the servicePolicy, and an error, if there is any.
func (c *FakeServicePolicies) Update(servicePolicy *v1alpha2.ServicePolicy) (result *v1alpha2.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicepoliciesResource, c.ns, servicePolicy), &v1alpha2.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ServicePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicePolicies) UpdateStatus(servicePolicy *v1alpha2.ServicePolicy) (*v1alpha2.ServicePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicepoliciesResource, "status", c.ns, servicePolicy), &v1alpha2.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ServicePolicy), err
}

// Delete takes name of the servicePolicy and deletes it. Returns an error if one occurs.
func (c *FakeServicePolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicepoliciesResource, c.ns, name), &v1alpha2.ServicePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicepoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.ServicePolicyList{})
	return err
}

// Patch applies the patch and returns the patched servicePolicy.
func (c *FakeServicePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha2.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ServicePolicy), err
}
