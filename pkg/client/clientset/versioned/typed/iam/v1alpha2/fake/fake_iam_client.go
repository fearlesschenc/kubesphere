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
	v1alpha2 "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/typed/iam/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIamV1alpha2 struct {
	*testing.Fake
}

func (c *FakeIamV1alpha2) GlobalRoles() v1alpha2.GlobalRoleInterface {
	return &FakeGlobalRoles{c}
}

func (c *FakeIamV1alpha2) GlobalRoleBindings() v1alpha2.GlobalRoleBindingInterface {
	return &FakeGlobalRoleBindings{c}
}

func (c *FakeIamV1alpha2) Groups() v1alpha2.GroupInterface {
	return &FakeGroups{c}
}

func (c *FakeIamV1alpha2) GroupBindings() v1alpha2.GroupBindingInterface {
	return &FakeGroupBindings{c}
}

func (c *FakeIamV1alpha2) LoginRecords() v1alpha2.LoginRecordInterface {
	return &FakeLoginRecords{c}
}

func (c *FakeIamV1alpha2) RoleBases() v1alpha2.RoleBaseInterface {
	return &FakeRoleBases{c}
}

func (c *FakeIamV1alpha2) Users() v1alpha2.UserInterface {
	return &FakeUsers{c}
}

func (c *FakeIamV1alpha2) WorkspaceRoles() v1alpha2.WorkspaceRoleInterface {
	return &FakeWorkspaceRoles{c}
}

func (c *FakeIamV1alpha2) WorkspaceRoleBindings() v1alpha2.WorkspaceRoleBindingInterface {
	return &FakeWorkspaceRoleBindings{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIamV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
