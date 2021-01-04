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
	"strings"
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

/* TODO: add openpitrix metrics
cond := &params.Conditions{
		Match: map[string]string{
			openpitrix.Status: openpitrix.StatusActive,
			openpitrix.RepoId: openpitrix.BuiltinRepoId,
		},
	}
	if h.op != nil {
		tmpl, err := h.op.ListApps(cond, "", false, 0, 0)
		if err != nil {
			res.Results = append(res.Results, monitoring.Metric{
				Name:  KubeSphereAppTmplCount,
				Error: err.Error(),
			})
		} else {
			res.Results = append(res.Results, monitoring.Metric{
				Name: KubeSphereAppTmplCount,
				MetricData: monitoring.MetricData{
					MetricType: monitoring.MetricTypeVector,
					MetricValues: []monitoring.MetricValue{
						{
							Sample: &monitoring.Point{now, float64(tmpl.TotalCount)},
						},
					},
				},
			})
		}
	}
*/

func (h handler) getKubeSphereMetrics(req *restful.Request, resp *restful.Response) {
	q := req.Request.URL.Query()
	q.Set(MetricsFilterQueryKey, strings.Join([]string{
		"kubesphere_workspace_count",
		"kubesphere_user_count",
		"kubesphere_cluser_count",
		"kubesphere_app_template_count",
	}, MetricsSep))
	q.Del(RangeStartQueryKey)
	q.Del(RangeEndQueryKey)
	q.Del(RangeStepQueryKey)
	req.Request.URL.RawQuery = q.Encode()

	h.getMetricsForObject(req, resp, h.monitoring.Cluster())
}

func (h handler) getClusterMetrics(req *restful.Request, resp *restful.Response) {
	h.getMetricsForObject(req, resp, h.monitoring.Cluster())
}

func (h handler) getNodeMetrics(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("node")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Node(name, filter))
}

func (h handler) getWorkspaceMetrics(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("workspace")
	filter := parseResourcesFilter(req)

	if req.QueryParameter("type") == "statistics" {
		req.Request.Form.Set(MetricsFilterQueryKey, strings.Join([]string{
			"workspace_namespace_count",
			"workspace_devops_project_count",
			"workspace_member_count",
			"workspace_role_count",
		}, MetricsSep))
	}

	h.getMetricsForObject(req, resp, h.monitoring.Workspace(name, filter))
}

func (h handler) getNamespaceMetrics(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("namespace")
	workspace := req.PathParameter("workspace")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Namespace(workspace, name, filter))
}

func (h handler) getWorkloadMetrics(req *restful.Request, resp *restful.Response) {
	kind := req.PathParameter("kind")
	namespace := req.PathParameter("namespace")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Workload(namespace, kind, filter))
}

func (h handler) getPodMetrics(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("pod")
	workloadKind := req.PathParameter("kind")
	workload := req.PathParameter("workload")
	namespace := req.PathParameter("namespace")
	node := req.PathParameter("node")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Pod(node, namespace, workloadKind, workload, name, filter))
}

func (h handler) getContainerMetrics(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("container")
	namespace := req.PathParameter("namespace")
	podName := req.PathParameter("pod")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.Container(namespace, podName, name, filter))
}

func (h handler) getPVCMetrics(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("pvc")
	storageClass := req.PathParameter("storageclass")
	namespace := req.PathParameter("namespace")
	filter := parseResourcesFilter(req)

	h.getMetricsForObject(req, resp, h.monitoring.PVC(namespace, storageClass, name, filter))
}

func (h handler) getComponentMetrics(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("component")

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
