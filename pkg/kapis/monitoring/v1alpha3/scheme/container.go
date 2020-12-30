package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var ContainerScheme = MetricScheme{
	"container_cpu_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Container).CPUUsage()
	},
	"container_memory_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Container).MemoryUsage()
	},
}
