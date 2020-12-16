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

package v1beta1

import (
	rest "k8s.io/client-go/rest"
	v1beta1 "github.com/fearlesschenc/kubesphere/pkg/apis/types/v1beta1"
	"github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned/scheme"
)

type TypesV1beta1Interface interface {
	RESTClient() rest.Interface
	FederatedApplicationsGetter
	FederatedClusterRolesGetter
	FederatedClusterRoleBindingsGetter
	FederatedConfigMapsGetter
	FederatedDeploymentsGetter
	FederatedIngressesGetter
	FederatedJobsGetter
	FederatedLimitRangesGetter
	FederatedNamespacesGetter
	FederatedPersistentVolumeClaimsGetter
	FederatedResourceQuotasGetter
	FederatedSecretsGetter
	FederatedServicesGetter
	FederatedStatefulSetsGetter
	FederatedUsersGetter
	FederatedWorkspacesGetter
}

// TypesV1beta1Client is used to interact with features provided by the types.kubefed.io group.
type TypesV1beta1Client struct {
	restClient rest.Interface
}

func (c *TypesV1beta1Client) FederatedApplications(namespace string) FederatedApplicationInterface {
	return newFederatedApplications(c, namespace)
}

func (c *TypesV1beta1Client) FederatedClusterRoles(namespace string) FederatedClusterRoleInterface {
	return newFederatedClusterRoles(c, namespace)
}

func (c *TypesV1beta1Client) FederatedClusterRoleBindings(namespace string) FederatedClusterRoleBindingInterface {
	return newFederatedClusterRoleBindings(c, namespace)
}

func (c *TypesV1beta1Client) FederatedConfigMaps(namespace string) FederatedConfigMapInterface {
	return newFederatedConfigMaps(c, namespace)
}

func (c *TypesV1beta1Client) FederatedDeployments(namespace string) FederatedDeploymentInterface {
	return newFederatedDeployments(c, namespace)
}

func (c *TypesV1beta1Client) FederatedIngresses(namespace string) FederatedIngressInterface {
	return newFederatedIngresses(c, namespace)
}

func (c *TypesV1beta1Client) FederatedJobs(namespace string) FederatedJobInterface {
	return newFederatedJobs(c, namespace)
}

func (c *TypesV1beta1Client) FederatedLimitRanges(namespace string) FederatedLimitRangeInterface {
	return newFederatedLimitRanges(c, namespace)
}

func (c *TypesV1beta1Client) FederatedNamespaces(namespace string) FederatedNamespaceInterface {
	return newFederatedNamespaces(c, namespace)
}

func (c *TypesV1beta1Client) FederatedPersistentVolumeClaims(namespace string) FederatedPersistentVolumeClaimInterface {
	return newFederatedPersistentVolumeClaims(c, namespace)
}

func (c *TypesV1beta1Client) FederatedResourceQuotas(namespace string) FederatedResourceQuotaInterface {
	return newFederatedResourceQuotas(c, namespace)
}

func (c *TypesV1beta1Client) FederatedSecrets(namespace string) FederatedSecretInterface {
	return newFederatedSecrets(c, namespace)
}

func (c *TypesV1beta1Client) FederatedServices(namespace string) FederatedServiceInterface {
	return newFederatedServices(c, namespace)
}

func (c *TypesV1beta1Client) FederatedStatefulSets(namespace string) FederatedStatefulSetInterface {
	return newFederatedStatefulSets(c, namespace)
}

func (c *TypesV1beta1Client) FederatedUsers(namespace string) FederatedUserInterface {
	return newFederatedUsers(c, namespace)
}

func (c *TypesV1beta1Client) FederatedWorkspaces() FederatedWorkspaceInterface {
	return newFederatedWorkspaces(c)
}

// NewForConfig creates a new TypesV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*TypesV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TypesV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new TypesV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TypesV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TypesV1beta1Client for the given RESTClient.
func New(c rest.Interface) *TypesV1beta1Client {
	return &TypesV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
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
func (c *TypesV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
