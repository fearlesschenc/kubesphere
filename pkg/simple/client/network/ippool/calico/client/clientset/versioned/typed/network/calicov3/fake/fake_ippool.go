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
	calicov3 "github.com/fearlesschenc/kubesphere/pkg/apis/network/calicov3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIPPools implements IPPoolInterface
type FakeIPPools struct {
	Fake *FakeCrdCalicov3
}

var ippoolsResource = schema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "calicov3", Resource: "ippools"}

var ippoolsKind = schema.GroupVersionKind{Group: "crd.projectcalico.org", Version: "calicov3", Kind: "IPPool"}

// Get takes name of the iPPool, and returns the corresponding iPPool object, and an error if there is any.
func (c *FakeIPPools) Get(name string, options v1.GetOptions) (result *calicov3.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ippoolsResource, name), &calicov3.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPPool), err
}

// List takes label and field selectors, and returns the list of IPPools that match those selectors.
func (c *FakeIPPools) List(opts v1.ListOptions) (result *calicov3.IPPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ippoolsResource, ippoolsKind, opts), &calicov3.IPPoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &calicov3.IPPoolList{ListMeta: obj.(*calicov3.IPPoolList).ListMeta}
	for _, item := range obj.(*calicov3.IPPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPPools.
func (c *FakeIPPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ippoolsResource, opts))
}

// Create takes the representation of a iPPool and creates it.  Returns the server's representation of the iPPool, and an error, if there is any.
func (c *FakeIPPools) Create(iPPool *calicov3.IPPool) (result *calicov3.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ippoolsResource, iPPool), &calicov3.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPPool), err
}

// Update takes the representation of a iPPool and updates it. Returns the server's representation of the iPPool, and an error, if there is any.
func (c *FakeIPPools) Update(iPPool *calicov3.IPPool) (result *calicov3.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ippoolsResource, iPPool), &calicov3.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPPool), err
}

// Delete takes name of the iPPool and deletes it. Returns an error if one occurs.
func (c *FakeIPPools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ippoolsResource, name), &calicov3.IPPool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ippoolsResource, listOptions)

	_, err := c.Fake.Invokes(action, &calicov3.IPPoolList{})
	return err
}

// Patch applies the patch and returns the patched iPPool.
func (c *FakeIPPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *calicov3.IPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ippoolsResource, name, pt, data, subresources...), &calicov3.IPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPPool), err
}
