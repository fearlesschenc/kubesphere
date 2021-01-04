package v1alpha3

import (
	"errors"
	"github.com/emicklei/go-restful"
	"github.com/fearlesschenc/kubesphere/pkg/api"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring"
	"github.com/prometheus-community/prom-label-proxy/injectproxy"
	prometheuslabels "github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql/parser"
)

/*
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
*/

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
