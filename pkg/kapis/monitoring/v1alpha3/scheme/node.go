package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var NodeScheme = MetricScheme{
	"node_cpu_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).CPUUtilization()
	},
	"node_cpu_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).CPUTotal()
	},
	"node_cpu_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).CPUUsage()
	},
	"node_memory_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).MemoryUtilization()
	},
	"node_memory_usage_wo_cache": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).MemoryUsageWoCache()
	},
	"node_memory_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).MemoryAvailable()
	},
	"node_memory_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).MemoryTotal()
	},
	"node_net_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).NetUtilization()
	},
	"node_net_bytes_transmitted": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).NetBytesTransmitted()
	},
	"node_net_bytes_received": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).NetBytesReceived()
	},
	"node_disk_read_iops": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskReadIOPS()
	},
	"node_disk_write_iops": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskWriteIOPS()
	},
	"node_disk_read_throughput": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskReadThroughput()
	},
	"node_disk_write_throughput": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskWriteThroughput()
	},
	"node_disk_size_capacity": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskSizeCapacity()
	},
	"node_disk_size_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskSizeAvailable()
	},
	"node_disk_size_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskSizeUsage()
	},
	"node_disk_size_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskSizeUtilization()
	},
	"node_disk_inode_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskInodeTotal()
	},
	"node_disk_inode_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskInodeUsage()
	},
	"node_disk_inode_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).DiskInodeUtilization()
	},
	"node_pod_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PodCount()
	},
	"node_pod_quota": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PodQuota()
	},
	"node_pod_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PodUtilization()
	},
	"node_pod_running_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PodRunning()
	},
	"node_pod_succeeded_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PodSucceed()
	},
	"node_pod_abnormal_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PodAbnormal()
	},
	"node_pod_abnormal_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PodAbnormalRatio()
	},
	"node_load1": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).Load1()
	},
	"node_load5": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).Load5()
	},
	"node_load15": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).Load15()
	},
	"node_pleg_quantile": func(obj types.Object) types.MetricQuery {
		return obj.(types.Node).PlegQuantile()
	},
}
