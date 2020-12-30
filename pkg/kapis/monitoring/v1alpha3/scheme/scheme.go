package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

type QueryForObject func(obj types.Object) types.MetricQuery

type MetricScheme map[string]QueryForObject

var EmptyScheme = MetricScheme{}

func (scheme MetricScheme) QueryForMetric(obj types.Object, metric string) types.MetricQuery {
	queryFunc, ok := scheme[metric]
	if !ok {
		return nil
	}

	return queryFunc(obj)
}

func (scheme MetricScheme) ListMetrics() (ret []string) {
	for k, _ := range scheme {
		ret = append(ret, k)
	}

	return ret
}

func GetScheme(obj types.Object) MetricScheme {
	switch obj.(type) {
	case types.Cluster:
		return ClusterScheme
	case types.Node:
		return NodeScheme
	case types.Workspace:
		return WorkloadScheme
	case types.Namespace:
		return NamespaceScheme
	case types.Workload:
		return WorkloadScheme
	case types.Pod:
		return PodScheme
	case types.Container:
		return ContainerScheme
	case types.PVC:
		return PVCScheme
	case types.Etcd:
		return EtcdScheme
	case types.APIServer:
		return APIServerScheme
	case types.Scheduler:
		return SchedulerScheme
	}

	return EmptyScheme
}

func QueryForMetric(obj types.Object, metric string) types.MetricQuery {
	return GetScheme(obj).QueryForMetric(obj, metric)
}
