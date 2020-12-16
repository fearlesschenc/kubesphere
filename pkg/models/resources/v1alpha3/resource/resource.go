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

package resource

import (
	"errors"

	snapshotv1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"github.com/fearlesschenc/kubesphere/pkg/api"
	devopsv1alpha3 "github.com/fearlesschenc/kubesphere/pkg/apis/devops/v1alpha3"
	iamv1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/iam/v1alpha2"
	tenantv1alpha1 "github.com/fearlesschenc/kubesphere/pkg/apis/tenant/v1alpha1"
	tenantv1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/tenant/v1alpha2"
	typesv1beta1 "github.com/fearlesschenc/kubesphere/pkg/apis/types/v1beta1"
	"github.com/fearlesschenc/kubesphere/pkg/apiserver/query"
	"github.com/fearlesschenc/kubesphere/pkg/informers"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/application"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/cluster"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/clusterrole"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/clusterrolebinding"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/configmap"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/customresourcedefinition"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/daemonset"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/deployment"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/devops"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatedapplication"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatedconfigmap"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federateddeployment"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatedingress"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatednamespace"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatedpersistentvolumeclaim"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatedsecret"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatedservice"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/federatedstatefulset"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/globalrole"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/globalrolebinding"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/group"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/groupbinding"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/ingress"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/job"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/loginrecord"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/namespace"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/networkpolicy"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/node"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/persistentvolumeclaim"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/pod"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/role"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/rolebinding"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/service"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/statefulset"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/user"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/volumesnapshot"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/workspace"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/workspacerole"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/workspacerolebinding"
	"github.com/fearlesschenc/kubesphere/pkg/models/resources/v1alpha3/workspacetemplate"
)

var ErrResourceNotSupported = errors.New("resource is not supported")

type ResourceGetter struct {
	getters map[schema.GroupVersionResource]v1alpha3.Interface
}

func NewResourceGetter(factory informers.InformerFactory) *ResourceGetter {
	getters := make(map[schema.GroupVersionResource]v1alpha3.Interface)

	getters[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}] = deployment.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "daemonsets"}] = daemonset.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "statefulsets"}] = statefulset.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "services"}] = service.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "namespaces"}] = namespace.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"}] = configmap.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}] = pod.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "nodes"}] = node.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "ingresses"}] = ingress.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "app.k8s.io", Version: "v1beta1", Resource: "applications"}] = application.New(factory.ApplicationSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "networkpolicies"}] = networkpolicy.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "batch", Version: "v1", Resource: "jobs"}] = job.New(factory.KubernetesSharedInformerFactory())

	// kubesphere resources
	getters[devopsv1alpha3.SchemeGroupVersion.WithResource(devopsv1alpha3.ResourcePluralDevOpsProject)] = devops.New(factory.KubeSphereSharedInformerFactory())
	getters[tenantv1alpha1.SchemeGroupVersion.WithResource(tenantv1alpha1.ResourcePluralWorkspace)] = workspace.New(factory.KubeSphereSharedInformerFactory())
	getters[tenantv1alpha1.SchemeGroupVersion.WithResource(tenantv1alpha2.ResourcePluralWorkspaceTemplate)] = workspacetemplate.New(factory.KubeSphereSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralGlobalRole)] = globalrole.New(factory.KubeSphereSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralWorkspaceRole)] = workspacerole.New(factory.KubeSphereSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralUser)] = user.New(factory.KubeSphereSharedInformerFactory(), factory.KubernetesSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralGlobalRoleBinding)] = globalrolebinding.New(factory.KubeSphereSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralWorkspaceRoleBinding)] = workspacerolebinding.New(factory.KubeSphereSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralLoginRecord)] = loginrecord.New(factory.KubeSphereSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcePluralGroup)] = group.New(factory.KubeSphereSharedInformerFactory())
	getters[iamv1alpha2.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcePluralGroupBinding)] = groupbinding.New(factory.KubeSphereSharedInformerFactory())
	getters[rbacv1.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralRole)] = role.New(factory.KubernetesSharedInformerFactory())
	getters[rbacv1.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralClusterRole)] = clusterrole.New(factory.KubernetesSharedInformerFactory())
	getters[rbacv1.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralRoleBinding)] = rolebinding.New(factory.KubernetesSharedInformerFactory())
	getters[rbacv1.SchemeGroupVersion.WithResource(iamv1alpha2.ResourcesPluralClusterRoleBinding)] = clusterrolebinding.New(factory.KubernetesSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "persistentvolumeclaims"}] = persistentvolumeclaim.New(factory.KubernetesSharedInformerFactory(), factory.SnapshotSharedInformerFactory())
	getters[snapshotv1beta1.SchemeGroupVersion.WithResource("volumesnapshots")] = volumesnapshot.New(factory.SnapshotSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "cluster.kubesphere.io", Version: "v1alpha1", Resource: "clusters"}] = cluster.New(factory.KubeSphereSharedInformerFactory())
	getters[schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}] = customresourcedefinition.New(factory.ApiExtensionSharedInformerFactory())

	// federated resources
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedNamespace)] = federatednamespace.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedDeployment)] = federateddeployment.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedSecret)] = federatedsecret.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedConfigmap)] = federatedconfigmap.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedService)] = federatedservice.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedApplication)] = federatedapplication.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedPersistentVolumeClaim)] = federatedpersistentvolumeclaim.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedStatefulSet)] = federatedstatefulset.New(factory.KubeSphereSharedInformerFactory())
	getters[typesv1beta1.SchemeGroupVersion.WithResource(typesv1beta1.ResourcePluralFederatedIngress)] = federatedingress.New(factory.KubeSphereSharedInformerFactory())

	return &ResourceGetter{
		getters: getters,
	}
}

// tryResource will retrieve a getter with resource name, it doesn't guarantee find resource with correct group version
// need to refactor this use schema.GroupVersionResource
func (r *ResourceGetter) tryResource(resource string) v1alpha3.Interface {
	for k, v := range r.getters {
		if k.Resource == resource {
			return v
		}
	}
	return nil
}

func (r *ResourceGetter) Get(resource, namespace, name string) (runtime.Object, error) {
	getter := r.tryResource(resource)
	if getter == nil {
		return nil, ErrResourceNotSupported
	}
	return getter.Get(namespace, name)
}

func (r *ResourceGetter) List(resource, namespace string, query *query.Query) (*api.ListResult, error) {
	getter := r.tryResource(resource)
	if getter == nil {
		return nil, ErrResourceNotSupported
	}
	return getter.List(namespace, query)
}
