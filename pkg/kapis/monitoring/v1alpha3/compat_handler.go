package v1alpha3

import (
	"errors"
	"github.com/emicklei/go-restful"
	"github.com/fearlesschenc/kubesphere/pkg/api"
	"github.com/fearlesschenc/kubesphere/pkg/constants"
	"github.com/fearlesschenc/kubesphere/pkg/models/openpitrix"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring"
	"github.com/fearlesschenc/kubesphere/pkg/server/params"
	"github.com/prometheus-community/prom-label-proxy/injectproxy"
	prometheuslabels "github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql/parser"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"time"
)

const (
	KubeSphereWorkspaceCount = "kubesphere_workspace_count"
	KubeSphereUserCount      = "kubesphere_user_count"
	KubeSphereClusterCount   = "kubesphere_cluser_count"
	KubeSphereAppTmplCount   = "kubesphere_app_template_count"

	WorkspaceNamespaceCount = "workspace_namespace_count"
	WorkspaceDevopsCount    = "workspace_devops_project_count"
	WorkspaceMemberCount    = "workspace_member_count"
	WorkspaceRoleCount      = "workspace_role_count"
)

func handleNoHit(namedMetrics []string) Metrics {
	var res Metrics
	for _, metric := range namedMetrics {
		res.Results = append(res.Results, monitoring.Metric{
			Name:       metric,
			MetricData: monitoring.MetricData{},
		})
	}
	return res
}

func (h handler) getKubeSphereMetrics(req *restful.Request, resp *restful.Response) {
	var res Metrics
	now := float64(time.Now().Unix())

	clusterList, err := h.ks.Cluster().V1alpha1().Clusters().Lister().List(labels.Everything())
	clusterTotal := len(clusterList)
	if clusterTotal == 0 {
		clusterTotal = 1
	}
	if err != nil {
		res.Results = append(res.Results, monitoring.Metric{
			Name:  KubeSphereClusterCount,
			Error: err.Error(),
		})
	} else {
		res.Results = append(res.Results, monitoring.Metric{
			Name: KubeSphereClusterCount,
			MetricData: monitoring.MetricData{
				MetricType: monitoring.MetricTypeVector,
				MetricValues: []monitoring.MetricValue{
					{
						Sample: &monitoring.Point{now, float64(clusterTotal)},
					},
				},
			},
		})
	}

	wkList, err := h.ks.Tenant().V1alpha2().WorkspaceTemplates().Lister().List(labels.Everything())
	if err != nil {
		res.Results = append(res.Results, monitoring.Metric{
			Name:  KubeSphereWorkspaceCount,
			Error: err.Error(),
		})
	} else {
		res.Results = append(res.Results, monitoring.Metric{
			Name: KubeSphereWorkspaceCount,
			MetricData: monitoring.MetricData{
				MetricType: monitoring.MetricTypeVector,
				MetricValues: []monitoring.MetricValue{
					{
						Sample: &monitoring.Point{now, float64(len(wkList))},
					},
				},
			},
		})
	}

	usrList, err := h.ks.Iam().V1alpha2().Users().Lister().List(labels.Everything())
	if err != nil {
		res.Results = append(res.Results, monitoring.Metric{
			Name:  KubeSphereUserCount,
			Error: err.Error(),
		})
	} else {
		res.Results = append(res.Results, monitoring.Metric{
			Name: KubeSphereUserCount,
			MetricData: monitoring.MetricData{
				MetricType: monitoring.MetricTypeVector,
				MetricValues: []monitoring.MetricValue{
					{
						Sample: &monitoring.Point{now, float64(len(usrList))},
					},
				},
			},
		})
	}

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

	resp.WriteAsJson(res)
}

func (h handler) GetWorkspaceStats(workspace string) Metrics {
	var res Metrics
	now := float64(time.Now().Unix())

	selector := labels.SelectorFromSet(labels.Set{constants.WorkspaceLabelKey: workspace})
	opt := metav1.ListOptions{LabelSelector: selector.String()}

	nsList, err := h.kubernetesClient.CoreV1().Namespaces().List(opt)
	if err != nil {
		res.Results = append(res.Results, monitoring.Metric{
			Name:  WorkspaceNamespaceCount,
			Error: err.Error(),
		})
	} else {
		res.Results = append(res.Results, monitoring.Metric{
			Name: WorkspaceNamespaceCount,
			MetricData: monitoring.MetricData{
				MetricType: monitoring.MetricTypeVector,
				MetricValues: []monitoring.MetricValue{
					{
						Sample: &monitoring.Point{now, float64(len(nsList.Items))},
					},
				},
			},
		})
	}

	devopsList, err := h.ks.Devops().V1alpha3().DevOpsProjects().Lister().List(selector)
	if err != nil {
		res.Results = append(res.Results, monitoring.Metric{
			Name:  WorkspaceDevopsCount,
			Error: err.Error(),
		})
	} else {
		res.Results = append(res.Results, monitoring.Metric{
			Name: WorkspaceDevopsCount,
			MetricData: monitoring.MetricData{
				MetricType: monitoring.MetricTypeVector,
				MetricValues: []monitoring.MetricValue{
					{
						Sample: &monitoring.Point{now, float64(len(devopsList))},
					},
				},
			},
		})
	}

	memberList, err := h.ks.Iam().V1alpha2().WorkspaceRoleBindings().Lister().List(selector)
	if err != nil {
		res.Results = append(res.Results, monitoring.Metric{
			Name:  WorkspaceMemberCount,
			Error: err.Error(),
		})
	} else {
		res.Results = append(res.Results, monitoring.Metric{
			Name: WorkspaceMemberCount,
			MetricData: monitoring.MetricData{
				MetricType: monitoring.MetricTypeVector,
				MetricValues: []monitoring.MetricValue{
					{
						Sample: &monitoring.Point{now, float64(len(memberList))},
					},
				},
			},
		})
	}

	roleList, err := h.ks.Iam().V1alpha2().WorkspaceRoles().Lister().List(selector)
	if err != nil {
		res.Results = append(res.Results, monitoring.Metric{
			Name:  WorkspaceRoleCount,
			Error: err.Error(),
		})
	} else {
		res.Results = append(res.Results, monitoring.Metric{
			Name: WorkspaceRoleCount,
			MetricData: monitoring.MetricData{
				MetricType: monitoring.MetricTypeVector,
				MetricValues: []monitoring.MetricValue{
					{
						Sample: &monitoring.Point{now, float64(len(roleList))},
					},
				},
			},
		})
	}

	return res
}

func (h handler) getMetadata(req *restful.Request, resp *restful.Response) {
	namespace := req.PathParameter("namespace")
	resp.WriteAsJson(h.monitoring.GetMetadata(namespace))
}

func makeExpr(input, ns string) (string, error) {
	root, err := parser.ParseExpr(input)
	if err != nil {
		return "", err
	}

	err = injectproxy.NewEnforcer(&prometheuslabels.Matcher{
		Type:  prometheuslabels.MatchEqual,
		Name:  "namespace",
		Value: ns,
	}).EnforceNode(root)
	if err != nil {
		return "", err
	}

	return root.String(), nil
}

func (h handler) getMetricLabelSet(req *restful.Request, resp *restful.Response) {
	metric := req.QueryParameter("metric")
	namespace := req.QueryParameter("namespace")
	start := req.QueryParameter("start")
	end := req.QueryParameter("end")

	if metric == "" || start == "" || end == "" {
		api.HandleBadRequest(resp, nil, errors.New("required fields are missing: [metric, start, end]"))
		return
	}

	expr, err := makeExpr(metric, namespace)
	if err != nil {
		api.HandleBadRequest(resp, nil, err)
		return
	}

	tr, err := parseTimeRange(req)
	if err != nil {
		api.HandleBadRequest(resp, nil, err)
		return
	}

	resp.WriteAsJson(h.monitoring.GetMetricLabelSet(expr, tr.Range.Start, tr.Range.End))
}

func (h handler) adhocQuery(req *restful.Request, resp *restful.Response) {
	expression := req.QueryParameter("expr")
	namespace := req.QueryParameter("namespace")

	expr, err := makeExpr(expression, namespace)
	if err != nil {
		api.HandleBadRequest(resp, nil, err)
		return
	}

	metric, err := getMetrics(req, []monitoring.MetricQuery{h.monitoring.Query(expr)}, QueryTimeout)
	if err != nil {
		api.HandleBadRequest(resp, nil, err)
		return
	}

	resp.WriteAsJson(metric[0])
}
