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
	v1beta1 "github.com/fearlesschenc/kubesphere/pkg/apis/types/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedClusterRoles implements FederatedClusterRoleInterface
type FakeFederatedClusterRoles struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedclusterrolesResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedclusterroles"}

var federatedclusterrolesKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedClusterRole"}

// Get takes name of the federatedClusterRole, and returns the corresponding federatedClusterRole object, and an error if there is any.
func (c *FakeFederatedClusterRoles) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedclusterrolesResource, c.ns, name), &v1beta1.FederatedClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRole), err
}

// List takes label and field selectors, and returns the list of FederatedClusterRoles that match those selectors.
func (c *FakeFederatedClusterRoles) List(opts v1.ListOptions) (result *v1beta1.FederatedClusterRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedclusterrolesResource, federatedclusterrolesKind, c.ns, opts), &v1beta1.FederatedClusterRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedClusterRoleList{ListMeta: obj.(*v1beta1.FederatedClusterRoleList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedClusterRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedClusterRoles.
func (c *FakeFederatedClusterRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedclusterrolesResource, c.ns, opts))

}

// Create takes the representation of a federatedClusterRole and creates it.  Returns the server's representation of the federatedClusterRole, and an error, if there is any.
func (c *FakeFederatedClusterRoles) Create(federatedClusterRole *v1beta1.FederatedClusterRole) (result *v1beta1.FederatedClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedclusterrolesResource, c.ns, federatedClusterRole), &v1beta1.FederatedClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRole), err
}

// Update takes the representation of a federatedClusterRole and updates it. Returns the server's representation of the federatedClusterRole, and an error, if there is any.
func (c *FakeFederatedClusterRoles) Update(federatedClusterRole *v1beta1.FederatedClusterRole) (result *v1beta1.FederatedClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedclusterrolesResource, c.ns, federatedClusterRole), &v1beta1.FederatedClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedClusterRoles) UpdateStatus(federatedClusterRole *v1beta1.FederatedClusterRole) (*v1beta1.FederatedClusterRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedclusterrolesResource, "status", c.ns, federatedClusterRole), &v1beta1.FederatedClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRole), err
}

// Delete takes name of the federatedClusterRole and deletes it. Returns an error if one occurs.
func (c *FakeFederatedClusterRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedclusterrolesResource, c.ns, name), &v1beta1.FederatedClusterRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedClusterRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedclusterrolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedClusterRoleList{})
	return err
}

// Patch applies the patch and returns the patched federatedClusterRole.
func (c *FakeFederatedClusterRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedClusterRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedclusterrolesResource, c.ns, name, pt, data, subresources...), &v1beta1.FederatedClusterRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRole), err
}
