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
	v1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/iam/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRoleBases implements RoleBaseInterface
type FakeRoleBases struct {
	Fake *FakeIamV1alpha2
}

var rolebasesResource = schema.GroupVersionResource{Group: "iam.kubesphere.io", Version: "v1alpha2", Resource: "rolebases"}

var rolebasesKind = schema.GroupVersionKind{Group: "iam.kubesphere.io", Version: "v1alpha2", Kind: "RoleBase"}

// Get takes name of the roleBase, and returns the corresponding roleBase object, and an error if there is any.
func (c *FakeRoleBases) Get(name string, options v1.GetOptions) (result *v1alpha2.RoleBase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(rolebasesResource, name), &v1alpha2.RoleBase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoleBase), err
}

// List takes label and field selectors, and returns the list of RoleBases that match those selectors.
func (c *FakeRoleBases) List(opts v1.ListOptions) (result *v1alpha2.RoleBaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(rolebasesResource, rolebasesKind, opts), &v1alpha2.RoleBaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.RoleBaseList{ListMeta: obj.(*v1alpha2.RoleBaseList).ListMeta}
	for _, item := range obj.(*v1alpha2.RoleBaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleBases.
func (c *FakeRoleBases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(rolebasesResource, opts))
}

// Create takes the representation of a roleBase and creates it.  Returns the server's representation of the roleBase, and an error, if there is any.
func (c *FakeRoleBases) Create(roleBase *v1alpha2.RoleBase) (result *v1alpha2.RoleBase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(rolebasesResource, roleBase), &v1alpha2.RoleBase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoleBase), err
}

// Update takes the representation of a roleBase and updates it. Returns the server's representation of the roleBase, and an error, if there is any.
func (c *FakeRoleBases) Update(roleBase *v1alpha2.RoleBase) (result *v1alpha2.RoleBase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(rolebasesResource, roleBase), &v1alpha2.RoleBase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoleBase), err
}

// Delete takes name of the roleBase and deletes it. Returns an error if one occurs.
func (c *FakeRoleBases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(rolebasesResource, name), &v1alpha2.RoleBase{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoleBases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(rolebasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.RoleBaseList{})
	return err
}

// Patch applies the patch and returns the patched roleBase.
func (c *FakeRoleBases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.RoleBase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(rolebasesResource, name, pt, data, subresources...), &v1alpha2.RoleBase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoleBase), err
}
