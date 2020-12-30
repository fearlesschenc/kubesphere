package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var PodScheme = MetricScheme{
	"pod_cpu_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Pod).CPUUsage()
	},
	"pod_memory_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Pod).MemoryUsage()
	},
	"pod_net_bytes_transmitted": func(obj types.Object) types.MetricQuery {
		return obj.(types.Pod).NetBytesTransmitted()
	},
	"pod_net_bytes_received": func(obj types.Object) types.MetricQuery {
		return obj.(types.Pod).NetBytesReceived()
	},
}
