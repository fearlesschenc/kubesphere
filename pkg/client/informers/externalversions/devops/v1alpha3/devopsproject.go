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

package v1alpha3

import (
	time "time"

	devopsv1alpha3 "github.com/fearlesschenc/kubesphere/pkg/apis/devops/v1alpha3"
	versioned "github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fearlesschenc/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/fearlesschenc/kubesphere/pkg/client/listers/devops/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DevOpsProjectInformer provides access to a shared informer and lister for
// DevOpsProjects.
type DevOpsProjectInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.DevOpsProjectLister
}

type devOpsProjectInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDevOpsProjectInformer constructs a new informer for DevOpsProject type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDevOpsProjectInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDevOpsProjectInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDevOpsProjectInformer constructs a new informer for DevOpsProject type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDevOpsProjectInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevopsV1alpha3().DevOpsProjects().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevopsV1alpha3().DevOpsProjects().Watch(options)
			},
		},
		&devopsv1alpha3.DevOpsProject{},
		resyncPeriod,
		indexers,
	)
}

func (f *devOpsProjectInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDevOpsProjectInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *devOpsProjectInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&devopsv1alpha3.DevOpsProject{}, f.defaultInformer)
}

func (f *devOpsProjectInformer) Lister() v1alpha3.DevOpsProjectLister {
	return v1alpha3.NewDevOpsProjectLister(f.Informer().GetIndexer())
}
