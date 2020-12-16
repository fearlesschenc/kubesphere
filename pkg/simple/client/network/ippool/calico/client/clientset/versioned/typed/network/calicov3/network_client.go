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

package calicov3

import (
	calicov3 "github.com/fearlesschenc/kubesphere/pkg/apis/network/calicov3"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/network/ippool/calico/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CrdCalicov3Interface interface {
	RESTClient() rest.Interface
	BlockAffinitiesGetter
	IPAMBlocksGetter
	IPPoolsGetter
}

// CrdCalicov3Client is used to interact with features provided by the crd.projectcalico.org group.
type CrdCalicov3Client struct {
	restClient rest.Interface
}

func (c *CrdCalicov3Client) BlockAffinities() BlockAffinityInterface {
	return newBlockAffinities(c)
}

func (c *CrdCalicov3Client) IPAMBlocks() IPAMBlockInterface {
	return newIPAMBlocks(c)
}

func (c *CrdCalicov3Client) IPPools() IPPoolInterface {
	return newIPPools(c)
}

// NewForConfig creates a new CrdCalicov3Client for the given config.
func NewForConfig(c *rest.Config) (*CrdCalicov3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CrdCalicov3Client{client}, nil
}

// NewForConfigOrDie creates a new CrdCalicov3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CrdCalicov3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CrdCalicov3Client for the given RESTClient.
func New(c rest.Interface) *CrdCalicov3Client {
	return &CrdCalicov3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := calicov3.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CrdCalicov3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
