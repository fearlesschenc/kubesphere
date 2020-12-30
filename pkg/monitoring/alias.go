package monitoring

import "github.com/fearlesschenc/kubesphere/pkg/monitoring/types"

type Metric = types.Metric
type Metadata = types.Metadata
type MetricData = types.MetricData
type MetricValue = types.MetricValue
type Point = types.Point
type Range = types.Range
type MetricQuery = types.MetricQuery
type NamedMetricQuery = types.NamedMetricQuery
type Object = types.Object

var NewNamedMetricQuery = types.NewNamedMetricQuery

const (
	MetricTypeMatrix = types.MetricTypeMatrix
	MetricTypeVector = types.MetricTypeVector
)
