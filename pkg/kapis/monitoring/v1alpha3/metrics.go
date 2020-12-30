package v1alpha3

import (
	"github.com/emicklei/go-restful"
	"github.com/fearlesschenc/kubesphere/pkg/api"
	"github.com/fearlesschenc/kubesphere/pkg/kapis/monitoring/v1alpha3/scheme"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring"
	"k8s.io/klog"
	"strings"
	"time"
)

const (
	MetricsSep = "|"
	MetricsEnd = "$"
)

const (
	OrderAscending  = "asc"
	OrderDescending = "desc"

	DefaultStep  = 10 * time.Minute
	DefaultOrder = OrderDescending
	DefaultPage  = 1
	DefaultLimit = 5
)

const QueryTimeout = 10 * time.Second

type Metrics struct {
	Results     []monitoring.Metric `json:"results" description:"actual array of results"`
	CurrentPage int                 `json:"page,omitempty" description:"current page returned"`
	TotalPages  int                 `json:"total_page,omitempty" description:"total number of pages"`
	TotalItems  int                 `json:"total_item,omitempty" description:"page size"`
}

type Metadata struct {
	Data []monitoring.Metadata `json:"data" description:"actual array of results"`
}

type MetricLabelSet struct {
	Data []map[string]string `json:"data" description:"actual array of results"`
}

func getMetrics(request *restful.Request, queries []monitoring.MetricQuery, timeout time.Duration) (results []monitoring.Metric, err error) {
	var tr *timeRange
	tr, err = parseTimeRange(request)
	if err != nil {
		return
	}

	// one by one
	for index, query := range queries {
		var metric *monitoring.Metric
		if tr.isRange {
			metric = query.QueryRange(tr.Range)
		} else {
			metric = query.Query(tr.Time)
		}
		if metric.Error != "" {
			klog.Errorf("query [%d] failed: %s", index, metric.Error)
		}

		if q, ok := query.(monitoring.NamedMetricQuery); ok {
			metric.Name = q.GetName()
		}

		results = append(results, *metric)
	}

	//var wg sync.WaitGroup
	//var mutex sync.Mutex
	//for _, query := range queries {
	//	wg.Add(1)
	//
	//	go func(query monitoring.MetricQuery) {
	//		defer wg.Done()
	//
	//		var metric *monitoring.Metric
	//		if tr.isRange {
	//			metric = query.QueryRange(tr.Range)
	//		} else {
	//			metric = query.Query(tr.Time)
	//		}
	//
	//		mutex.Lock()
	//		defer mutex.Unlock()
	//		results = append(results, *metric)
	//	}(query)
	//}
	//
	//stop := make(chan struct{})
	//go func() {
	//	defer close(stop)
	//	wg.Wait()
	//}()
	//
	//select {
	//case <-stop:
	//	return
	//case <-time.After(timeout):
	//	return nil, fmt.Errorf("query timeout error")
	//}

	return
}

func parseMetricsFilter(request *restful.Request) string {
	pat := request.QueryParameter("metrics_filter")
	if pat == "" {
		pat = DefaultFilter
	}

	return pat
}

func (h handler) getMetricsForObject(request *restful.Request, response *restful.Response, obj monitoring.Object) {
	pattern := parseMetricsFilter(request)

	// improvement for current | separate metrics match
	var metrics []string
	if pattern == DefaultFilter {
		metrics = scheme.GetScheme(obj).ListMetrics()
	} else {
		metrics = strings.Split(strings.TrimSuffix(pattern, MetricsEnd), MetricsSep)
	}

	var queries []monitoring.MetricQuery
	for _, metric := range metrics {
		query := scheme.QueryForMetric(obj, metric)
		// ignore nil queryMetrics
		if query == nil {
			continue
		}

		queries = append(queries, monitoring.NewNamedMetricQuery(metric, query))
	}

	results, err := getMetrics(request, queries, QueryTimeout)
	if err != nil {
		api.HandleBadRequest(response, nil, err)
		return
	}

	ret := &Metrics{Results: results}
	response.WriteAsJson(ret)
	// TODO: sort
	return
}
