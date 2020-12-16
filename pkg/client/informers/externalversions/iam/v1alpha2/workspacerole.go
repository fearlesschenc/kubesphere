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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	iamv1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/iam/v1alpha2"
	versioned "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fearlesschenc/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/fearlesschenc/kubesphere/pkg/client/listers/iam/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WorkspaceRoleInformer provides access to a shared informer and lister for
// WorkspaceRoles.
type WorkspaceRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.WorkspaceRoleLister
}

type workspaceRoleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewWorkspaceRoleInformer constructs a new informer for WorkspaceRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkspaceRoleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkspaceRoleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredWorkspaceRoleInformer constructs a new informer for WorkspaceRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkspaceRoleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1alpha2().WorkspaceRoles().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1alpha2().WorkspaceRoles().Watch(options)
			},
		},
		&iamv1alpha2.WorkspaceRole{},
		resyncPeriod,
		indexers,
	)
}

func (f *workspaceRoleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkspaceRoleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workspaceRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&iamv1alpha2.WorkspaceRole{}, f.defaultInformer)
}

func (f *workspaceRoleInformer) Lister() v1alpha2.WorkspaceRoleLister {
	return v1alpha2.NewWorkspaceRoleLister(f.Informer().GetIndexer())
}
