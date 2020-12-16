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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/iam/v1alpha2"
)

// RoleBaseLister helps list RoleBases.
type RoleBaseLister interface {
	// List lists all RoleBases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.RoleBase, err error)
	// Get retrieves the RoleBase from the index for a given name.
	Get(name string) (*v1alpha2.RoleBase, error)
	RoleBaseListerExpansion
}

// roleBaseLister implements the RoleBaseLister interface.
type roleBaseLister struct {
	indexer cache.Indexer
}

// NewRoleBaseLister returns a new RoleBaseLister.
func NewRoleBaseLister(indexer cache.Indexer) RoleBaseLister {
	return &roleBaseLister{indexer: indexer}
}

// List lists all RoleBases in the indexer.
func (s *roleBaseLister) List(selector labels.Selector) (ret []*v1alpha2.RoleBase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.RoleBase))
	})
	return ret, err
}

// Get retrieves the RoleBase from the index for a given name.
func (s *roleBaseLister) Get(name string) (*v1alpha2.RoleBase, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("rolebase"), name)
	}
	return obj.(*v1alpha2.RoleBase), nil
}
