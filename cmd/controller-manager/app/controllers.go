/*
Copyright 2019 The KubeSphere Authors.

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

package app

import (
	iamv1alpha2 "github.com/fearlesschenc/kubesphere/pkg/apis/iam/v1alpha2"
	authoptions "github.com/fearlesschenc/kubesphere/pkg/apiserver/authentication/options"
	"github.com/fearlesschenc/kubesphere/pkg/controller/application"
	"github.com/fearlesschenc/kubesphere/pkg/controller/certificatesigningrequest"
	"github.com/fearlesschenc/kubesphere/pkg/controller/cluster"
	"github.com/fearlesschenc/kubesphere/pkg/controller/clusterrolebinding"
	"github.com/fearlesschenc/kubesphere/pkg/controller/destinationrule"
	"github.com/fearlesschenc/kubesphere/pkg/controller/devopscredential"
	"github.com/fearlesschenc/kubesphere/pkg/controller/devopsproject"
	"github.com/fearlesschenc/kubesphere/pkg/controller/globalrole"
	"github.com/fearlesschenc/kubesphere/pkg/controller/globalrolebinding"
	"github.com/fearlesschenc/kubesphere/pkg/controller/group"
	"github.com/fearlesschenc/kubesphere/pkg/controller/groupbinding"
	"github.com/fearlesschenc/kubesphere/pkg/controller/job"
	"github.com/fearlesschenc/kubesphere/pkg/controller/network/ippool"
	"github.com/fearlesschenc/kubesphere/pkg/controller/network/nsnetworkpolicy"
	"github.com/fearlesschenc/kubesphere/pkg/controller/network/nsnetworkpolicy/provider"
	"github.com/fearlesschenc/kubesphere/pkg/controller/pipeline"
	"github.com/fearlesschenc/kubesphere/pkg/controller/s2ibinary"
	"github.com/fearlesschenc/kubesphere/pkg/controller/s2irun"
	"github.com/fearlesschenc/kubesphere/pkg/controller/storage/capability"
	"github.com/fearlesschenc/kubesphere/pkg/controller/storage/expansion"
	"github.com/fearlesschenc/kubesphere/pkg/controller/user"
	"github.com/fearlesschenc/kubesphere/pkg/controller/virtualservice"
	"github.com/fearlesschenc/kubesphere/pkg/controller/workspacerole"
	"github.com/fearlesschenc/kubesphere/pkg/controller/workspacerolebinding"
	"github.com/fearlesschenc/kubesphere/pkg/controller/workspacetemplate"
	"github.com/fearlesschenc/kubesphere/pkg/informers"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/devops"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/k8s"
	ldapclient "github.com/fearlesschenc/kubesphere/pkg/simple/client/ldap"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/network"
	ippoolclient "github.com/fearlesschenc/kubesphere/pkg/simple/client/network/ippool"
	calicoclient "github.com/fearlesschenc/kubesphere/pkg/simple/client/network/ippool/calico"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/openpitrix"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/s3"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/kubefed/pkg/controller/util"
)

func addControllers(
	mgr manager.Manager,
	client k8s.Client,
	informerFactory informers.InformerFactory,
	devopsClient devops.Interface,
	s3Client s3.Interface,
	ldapClient ldapclient.Interface,
	options *k8s.KubernetesOptions,
	authenticationOptions *authoptions.AuthenticationOptions,
	openpitrixClient openpitrix.Client,
	multiClusterEnabled bool,
	networkOptions *network.Options,
	serviceMeshEnabled bool,
	kubectlImage string,
	stopCh <-chan struct{}) error {

	kubernetesInformer := informerFactory.KubernetesSharedInformerFactory()
	istioInformer := informerFactory.IstioSharedInformerFactory()
	kubesphereInformer := informerFactory.KubeSphereSharedInformerFactory()
	applicationInformer := informerFactory.ApplicationSharedInformerFactory()

	var vsController, drController manager.Runnable
	if serviceMeshEnabled {
		vsController = virtualservice.NewVirtualServiceController(kubernetesInformer.Core().V1().Services(),
			istioInformer.Networking().V1alpha3().VirtualServices(),
			istioInformer.Networking().V1alpha3().DestinationRules(),
			kubesphereInformer.Servicemesh().V1alpha2().Strategies(),
			client.Kubernetes(),
			client.Istio(),
			client.KubeSphere())

		drController = destinationrule.NewDestinationRuleController(kubernetesInformer.Apps().V1().Deployments(),
			istioInformer.Networking().V1alpha3().DestinationRules(),
			kubernetesInformer.Core().V1().Services(),
			kubesphereInformer.Servicemesh().V1alpha2().ServicePolicies(),
			client.Kubernetes(),
			client.Istio(),
			client.KubeSphere())
	}

	apController := application.NewApplicationController(kubernetesInformer.Core().V1().Services(),
		kubernetesInformer.Apps().V1().Deployments(),
		kubernetesInformer.Apps().V1().StatefulSets(),
		kubesphereInformer.Servicemesh().V1alpha2().Strategies(),
		kubesphereInformer.Servicemesh().V1alpha2().ServicePolicies(),
		applicationInformer.App().V1beta1().Applications(),
		client.Kubernetes(),
		client.Application())

	jobController := job.NewJobController(kubernetesInformer.Batch().V1().Jobs(), client.Kubernetes())

	var s2iBinaryController, s2iRunController, devopsProjectController, devopsPipelineController, devopsCredentialController manager.Runnable
	if devopsClient != nil {
		s2iBinaryController = s2ibinary.NewController(client.Kubernetes(),
			client.KubeSphere(),
			kubesphereInformer.Devops().V1alpha1().S2iBinaries(),
			s3Client,
		)

		s2iRunController = s2irun.NewS2iRunController(client.Kubernetes(),
			client.KubeSphere(),
			kubesphereInformer.Devops().V1alpha1().S2iBinaries(),
			kubesphereInformer.Devops().V1alpha1().S2iRuns())

		devopsProjectController = devopsproject.NewController(client.Kubernetes(),
			client.KubeSphere(), devopsClient,
			informerFactory.KubernetesSharedInformerFactory().Core().V1().Namespaces(),
			informerFactory.KubeSphereSharedInformerFactory().Devops().V1alpha3().DevOpsProjects(),
			informerFactory.KubeSphereSharedInformerFactory().Tenant().V1alpha1().Workspaces())

		devopsPipelineController = pipeline.NewController(client.Kubernetes(),
			client.KubeSphere(),
			devopsClient,
			informerFactory.KubernetesSharedInformerFactory().Core().V1().Namespaces(),
			informerFactory.KubeSphereSharedInformerFactory().Devops().V1alpha3().Pipelines())

		devopsCredentialController = devopscredential.NewController(client.Kubernetes(),
			devopsClient,
			informerFactory.KubernetesSharedInformerFactory().Core().V1().Namespaces(),
			informerFactory.KubernetesSharedInformerFactory().Core().V1().Secrets())

	}

	storageCapabilityController := capability.NewController(
		client.KubeSphere().StorageV1alpha1().StorageClassCapabilities(),
		kubesphereInformer.Storage().V1alpha1(),
		client.Kubernetes().StorageV1().StorageClasses(),
		kubernetesInformer.Storage().V1().StorageClasses(),
		capability.SnapshotSupported(client.Kubernetes().Discovery()),
		client.Snapshot().SnapshotV1beta1().VolumeSnapshotClasses(),
		informerFactory.SnapshotSharedInformerFactory().Snapshot().V1beta1().VolumeSnapshotClasses(),
		kubernetesInformer.Storage().V1beta1().CSIDrivers(),
	)

	volumeExpansionController := expansion.NewVolumeExpansionController(
		client.Kubernetes(),
		kubernetesInformer.Core().V1().PersistentVolumeClaims(),
		kubernetesInformer.Storage().V1().StorageClasses(),
		kubernetesInformer.Core().V1().Pods(),
		kubernetesInformer.Apps().V1().Deployments(),
		kubernetesInformer.Apps().V1().ReplicaSets(),
		kubernetesInformer.Apps().V1().StatefulSets())

	var fedUserCache, fedGlobalRoleBindingCache, fedGlobalRoleCache,
		fedWorkspaceRoleCache, fedWorkspaceRoleBindingCache cache.Store
	var fedUserCacheController, fedGlobalRoleBindingCacheController, fedGlobalRoleCacheController,
		fedWorkspaceRoleCacheController, fedWorkspaceRoleBindingCacheController cache.Controller

	if multiClusterEnabled {
		fedUserClient, err := util.NewResourceClient(client.Config(), &iamv1alpha2.FedUserResource)
		if err != nil {
			klog.Error(err)
			return err
		}
		fedGlobalRoleClient, err := util.NewResourceClient(client.Config(), &iamv1alpha2.FedGlobalRoleResource)
		if err != nil {
			klog.Error(err)
			return err
		}
		fedGlobalRoleBindingClient, err := util.NewResourceClient(client.Config(), &iamv1alpha2.FedGlobalRoleBindingResource)
		if err != nil {
			klog.Error(err)
			return err
		}
		fedWorkspaceRoleClient, err := util.NewResourceClient(client.Config(), &iamv1alpha2.FedWorkspaceRoleResource)
		if err != nil {
			klog.Error(err)
			return err
		}
		fedWorkspaceRoleBindingClient, err := util.NewResourceClient(client.Config(), &iamv1alpha2.FedWorkspaceRoleBindingResource)
		if err != nil {
			klog.Error(err)
			return err
		}

		fedUserCache, fedUserCacheController = util.NewResourceInformer(fedUserClient, "", &iamv1alpha2.FedUserResource, func(object runtime.Object) {})
		fedGlobalRoleCache, fedGlobalRoleCacheController = util.NewResourceInformer(fedGlobalRoleClient, "", &iamv1alpha2.FedGlobalRoleResource, func(object runtime.Object) {})
		fedGlobalRoleBindingCache, fedGlobalRoleBindingCacheController = util.NewResourceInformer(fedGlobalRoleBindingClient, "", &iamv1alpha2.FedGlobalRoleBindingResource, func(object runtime.Object) {})
		fedWorkspaceRoleCache, fedWorkspaceRoleCacheController = util.NewResourceInformer(fedWorkspaceRoleClient, "", &iamv1alpha2.FedWorkspaceRoleResource, func(object runtime.Object) {})
		fedWorkspaceRoleBindingCache, fedWorkspaceRoleBindingCacheController = util.NewResourceInformer(fedWorkspaceRoleBindingClient, "", &iamv1alpha2.FedWorkspaceRoleBindingResource, func(object runtime.Object) {})

		go fedUserCacheController.Run(stopCh)
		go fedGlobalRoleCacheController.Run(stopCh)
		go fedGlobalRoleBindingCacheController.Run(stopCh)
		go fedWorkspaceRoleCacheController.Run(stopCh)
		go fedWorkspaceRoleBindingCacheController.Run(stopCh)
	}

	userController := user.NewUserController(client.Kubernetes(), client.KubeSphere(),
		client.Config(),
		kubesphereInformer.Iam().V1alpha2().Users(),
		fedUserCache, fedUserCacheController,
		kubesphereInformer.Iam().V1alpha2().LoginRecords(),
		kubernetesInformer.Core().V1().ConfigMaps(),
		ldapClient, devopsClient,
		authenticationOptions, multiClusterEnabled)

	loginRecordController := user.NewLoginRecordController(
		client.Kubernetes(),
		client.KubeSphere(),
		kubesphereInformer.Iam().V1alpha2().LoginRecords(),
		authenticationOptions.LoginHistoryRetentionPeriod)

	csrController := certificatesigningrequest.NewController(client.Kubernetes(),
		kubernetesInformer.Certificates().V1beta1().CertificateSigningRequests(),
		kubernetesInformer.Core().V1().ConfigMaps(), client.Config())

	clusterRoleBindingController := clusterrolebinding.NewController(client.Kubernetes(),
		kubernetesInformer.Rbac().V1().ClusterRoleBindings(),
		kubernetesInformer.Apps().V1().Deployments(),
		kubernetesInformer.Core().V1().Pods(),
		kubesphereInformer.Iam().V1alpha2().Users(),
		kubectlImage)

	globalRoleController := globalrole.NewController(client.Kubernetes(), client.KubeSphere(),
		kubesphereInformer.Iam().V1alpha2().GlobalRoles(), fedGlobalRoleCache, fedGlobalRoleCacheController)

	workspaceRoleController := workspacerole.NewController(client.Kubernetes(), client.KubeSphere(),
		kubesphereInformer.Iam().V1alpha2().WorkspaceRoles(),
		fedWorkspaceRoleCache, fedWorkspaceRoleCacheController,
		kubesphereInformer.Tenant().V1alpha2().WorkspaceTemplates(), multiClusterEnabled)

	globalRoleBindingController := globalrolebinding.NewController(client.Kubernetes(), client.KubeSphere(),
		kubesphereInformer.Iam().V1alpha2().GlobalRoleBindings(),
		fedGlobalRoleBindingCache, fedGlobalRoleBindingCacheController,
		multiClusterEnabled)

	workspaceRoleBindingController := workspacerolebinding.NewController(client.Kubernetes(), client.KubeSphere(),
		kubesphereInformer.Iam().V1alpha2().WorkspaceRoleBindings(),
		fedWorkspaceRoleBindingCache, fedWorkspaceRoleBindingCacheController,
		kubesphereInformer.Tenant().V1alpha2().WorkspaceTemplates(), multiClusterEnabled)

	workspaceTemplateController := workspacetemplate.NewController(client.Kubernetes(), client.KubeSphere(),
		kubesphereInformer.Tenant().V1alpha2().WorkspaceTemplates(),
		kubesphereInformer.Tenant().V1alpha1().Workspaces(),
		kubesphereInformer.Iam().V1alpha2().RoleBases(),
		kubesphereInformer.Iam().V1alpha2().WorkspaceRoles(),
		kubesphereInformer.Types().V1beta1().FederatedWorkspaces(),
		multiClusterEnabled)

	groupBindingController := groupbinding.NewController(client.Kubernetes(), client.KubeSphere(),
		kubesphereInformer.Iam().V1alpha2().GroupBindings())

	groupController := group.NewController(client.Kubernetes(), client.KubeSphere(),
		kubesphereInformer.Iam().V1alpha2().Groups())

	var clusterController manager.Runnable
	if multiClusterEnabled {
		clusterController = cluster.NewClusterController(
			client.Kubernetes(),
			client.Config(),
			kubesphereInformer.Cluster().V1alpha1().Clusters(),
			client.KubeSphere().ClusterV1alpha1().Clusters(),
			openpitrixClient)
	}

	var nsnpController manager.Runnable
	if networkOptions.EnableNetworkPolicy {
		nsnpProvider, err := provider.NewNsNetworkPolicyProvider(client.Kubernetes(), kubernetesInformer.Networking().V1().NetworkPolicies())
		if err != nil {
			return err
		}

		nsnpController = nsnetworkpolicy.NewNSNetworkPolicyController(client.Kubernetes(),
			client.KubeSphere().NetworkV1alpha1(),
			kubesphereInformer.Network().V1alpha1().NamespaceNetworkPolicies(),
			kubernetesInformer.Core().V1().Services(),
			kubernetesInformer.Core().V1().Nodes(),
			kubesphereInformer.Tenant().V1alpha1().Workspaces(),
			kubernetesInformer.Core().V1().Namespaces(), nsnpProvider, networkOptions.NSNPOptions)
	}

	var ippoolController manager.Runnable
	if networkOptions.EnableIPPool {
		var ippoolProvider ippoolclient.Provider
		ippoolProvider = ippoolclient.NewProvider(client.KubeSphere(), networkOptions.IPPoolOptions)
		if networkOptions.IPPoolOptions.Calico != nil {
			ippoolProvider = calicoclient.NewProvider(client.KubeSphere(), *networkOptions.IPPoolOptions.Calico, options)
		}
		ippoolController = ippool.NewIPPoolController(kubesphereInformer.Network().V1alpha1().IPPools(),
			kubesphereInformer.Network().V1alpha1().IPAMBlocks(),
			client.Kubernetes(),
			client.KubeSphere(),
			networkOptions.IPPoolOptions,
			ippoolProvider)
	}

	controllers := map[string]manager.Runnable{
		"virtualservice-controller":       vsController,
		"destinationrule-controller":      drController,
		"application-controller":          apController,
		"job-controller":                  jobController,
		"s2ibinary-controller":            s2iBinaryController,
		"s2irun-controller":               s2iRunController,
		"storagecapability-controller":    storageCapabilityController,
		"volumeexpansion-controller":      volumeExpansionController,
		"user-controller":                 userController,
		"loginrecord-controller":          loginRecordController,
		"cluster-controller":              clusterController,
		"nsnp-controller":                 nsnpController,
		"csr-controller":                  csrController,
		"clusterrolebinding-controller":   clusterRoleBindingController,
		"globalrolebinding-controller":    globalRoleBindingController,
		"workspacetemplate-controller":    workspaceTemplateController,
		"workspacerole-controller":        workspaceRoleController,
		"workspacerolebinding-controller": workspaceRoleBindingController,
		"ippool-controller":               ippoolController,
		"groupbinding-controller":         groupBindingController,
		"group-controller":                groupController,
	}

	if devopsClient != nil {
		controllers["pipeline-controller"] = devopsPipelineController
		controllers["devopsprojects-controller"] = devopsProjectController
		controllers["devopscredential-controller"] = devopsCredentialController
	}

	if multiClusterEnabled {
		controllers["globalrole-controller"] = globalRoleController
	}

	for name, ctrl := range controllers {
		if ctrl == nil {
			klog.V(4).Infof("%s is not going to run due to dependent component disabled.", name)
			continue
		}

		if err := mgr.Add(ctrl); err != nil {
			klog.Error(err, "add controller to manager failed", "name", name)
			return err
		}
	}

	return nil
}
