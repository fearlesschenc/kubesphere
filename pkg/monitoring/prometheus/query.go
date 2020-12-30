package prometheus

import (
	"context"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prometheusmodel "github.com/prometheus/common/model"
	"time"
)

type Query struct {
	query string
	api   prometheusapiv1.API
}

func (q *Query) Query(ts time.Time) *types.Metric {
	var metric types.Metric

	// TODO: handle api.warnings
	val, _, err := q.api.Query(context.Background(), q.query, ts)
	if err != nil {
		metric.Error = err.Error()
		return &metric
	}

	metric.MetricData = parseQueryResp(val)
	return &metric
}

func (q *Query) QueryRange(r types.Range) *types.Metric {
	var metric types.Metric

	val, _, err := q.api.QueryRange(context.Background(), q.query, prometheusapiv1.Range{
		Start: r.Start,
		End:   r.End,
		Step:  r.Step,
	})
	if err != nil {
		metric.Error = err.Error()
		return &metric
	}

	metric.MetricData = parseQueryRangeResp(val)
	return &metric
}

func parseQueryResp(value prometheusmodel.Value) types.MetricData {
	res := types.MetricData{MetricType: types.MetricTypeVector}

	data, _ := value.(prometheusmodel.Vector)
	for _, v := range data {
		mv := types.MetricValue{
			Metadata: make(map[string]string),
		}

		for k, v := range v.Metric {
			mv.Metadata[string(k)] = string(v)
		}

		mv.Sample = &types.Point{float64(v.Timestamp) / 1000, float64(v.Value)}
		res.MetricValues = append(res.MetricValues, mv)
	}

	return res
}

func parseQueryRangeResp(value prometheusmodel.Value) types.MetricData {
	res := types.MetricData{MetricType: types.MetricTypeMatrix}

	data, _ := value.(prometheusmodel.Matrix)

	for _, v := range data {
		mv := types.MetricValue{
			Metadata: make(map[string]string),
		}

		for k, v := range v.Metric {
			mv.Metadata[string(k)] = string(v)
		}

		for _, k := range v.Values {
			mv.Series = append(mv.Series, types.Point{float64(k.Timestamp) / 1000, float64(k.Value)})
		}

		res.MetricValues = append(res.MetricValues, mv)
	}

	return res
}
