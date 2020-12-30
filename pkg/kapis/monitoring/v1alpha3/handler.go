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

package v1alpha3

import (
	"github.com/emicklei/go-restful"
	ksinformers "github.com/fearlesschenc/kubesphere/pkg/client/informers/externalversions"
	"github.com/fearlesschenc/kubesphere/pkg/informers"
	"github.com/fearlesschenc/kubesphere/pkg/models/openpitrix"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring"
	opclient "github.com/fearlesschenc/kubesphere/pkg/simple/client/openpitrix"
	"k8s.io/client-go/kubernetes"
)

const DefaultFilter = ".*"

type handler struct {
	kubernetesClient kubernetes.Interface
	monitoring       monitoring.Interface
	ks               ksinformers.SharedInformerFactory
	op               openpitrix.Interface
}

func newHandler(k kubernetes.Interface, m monitoring.Interface, f informers.InformerFactory, o opclient.Client) *handler {
	return &handler{
		kubernetesClient: k,
		monitoring:       m,
		ks:               f.KubeSphereSharedInformerFactory(),
		op:               openpitrix.NewOpenpitrixOperator(f.KubernetesSharedInformerFactory(), o),
	}
}

func parseResourcesFilter(req *restful.Request) string {
	filter := req.QueryParameter("resources_filter")
	if filter == "" {
		return DefaultFilter
	}

	return filter
}

func (h handler) getClusterMetrics(req *restful.Request, resp *restful.Response) {
	h.getMetricsForObject(req, resp, h.monitoring.Cluster())
}

func (h handler) getNodeMetrics(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("node")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Node(name, filter))
}

func (h handler) getWorkspaceMetrics(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("workspace")
	filter := parseResourcesFilter(req)

	typ := req.QueryParameter("type")
	if typ == "statistics" {
		h.GetWorkspaceStats(name)
		return
	}

	h.getMetricsForObject(req, resp, h.monitoring.Workspace(name, filter))
}

func (h handler) getNamespaceMetrics(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("namespace")
	workspace := req.QueryParameter("workspace")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Namespace(workspace, name, filter))
}

func (h handler) getWorkloadMetrics(req *restful.Request, resp *restful.Response) {
	kind := req.QueryParameter("kind")
	namespace := req.QueryParameter("namespace")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Workload(namespace, kind, filter))
}

func (h handler) getPodMetrics(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("pod")
	workloadKind := req.QueryParameter("kind")
	workload := req.QueryParameter("workload")
	namespace := req.QueryParameter("namespace")
	node := req.QueryParameter("node")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Pod(node, namespace, workloadKind, workload, name, filter))
}

func (h handler) getContainerMetrics(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("container")
	namespace := req.QueryParameter("namespace")
	podName := req.QueryParameter("pod")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Container(namespace, podName, name, filter))
}

func (h handler) getPVCMetrics(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("pvc")
	storageClass := req.QueryParameter("storageclass")
	namespace := req.QueryParameter("namespace")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.PVC(namespace, storageClass, name, filter))
}

func (h handler) getComponentMetrics(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("name")

	var component monitoring.Object
	switch name {
	case "etcd":
		component = h.monitoring.Component().Scheduler()
	case "apiserver":
		component = h.monitoring.Component().APIServer()
	case "scheduler":
		component = h.monitoring.Component().Scheduler()
	}

	h.getMetricsForObject(req, resp, component)
}
