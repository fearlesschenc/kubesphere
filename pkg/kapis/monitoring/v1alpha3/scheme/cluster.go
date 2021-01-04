package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var ClusterScheme = MetricScheme{
	"cluster_cpu_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).CPUUtilization()
	},
	"cluster_cpu_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).CPUUsage()
	},
	"cluster_cpu_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).CPUTotal()
	},
	"cluster_memory_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).MemoryUtilization()
	},
	"cluster_memory_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).MemoryAvailable()
	},
	"cluster_memory_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).MemoryTotal()
	},
	"cluster_memory_usage_wo_cache": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).MemoryUsageWoCache()
	},
	"cluster_net_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NetUtilization()
	},
	"cluster_net_bytes_transmitted": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NetBytesTransmitted()
	},
	"cluster_net_bytes_received": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NetBytesReceived()
	},
	"cluster_disk_read_iops": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskReadIOPS()
	},
	"cluster_disk_write_iops": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskWriteIOPS()
	},
	"cluster_disk_read_throughput": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskReadThroughput()
	},
	"cluster_disk_write_throughput": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskWriteThroughput()
	},
	"cluster_disk_size_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskSizeUsage()
	},
	"cluster_disk_size_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskSizeUtilization()
	},
	"cluster_disk_size_capacity": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskSizeCapacity()
	},
	"cluster_disk_size_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskSizeAvailable()
	},
	"cluster_disk_inode_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskInodeTotal()
	},
	"cluster_disk_inode_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskInodeUsage()
	},
	"cluster_disk_inode_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DiskInodeUtilization()
	},
	"cluster_namespace_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NamespaceCount()
	},
	"cluster_pod_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PodCount()
	},
	"cluster_pod_quota": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PodQuota()
	},
	"cluster_pod_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PodUtilization()
	},
	"cluster_pod_running_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PodRunning()
	},
	"cluster_pod_succeeded_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PodSucceed()
	},
	"cluster_pod_abnormal_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PodAbnormal()
	},
	"cluster_pod_abnormal_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PodAbnormalRatio()
	},
	"cluster_node_online": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NodeOnline()
	},
	"cluster_node_offline": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NodeOffline()
	},
	"cluster_node_offline_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NodeOfflineRatio()
	},
	"cluster_node_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).NodeTotal()
	},
	"cluster_cronjob_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).CronJobCount()
	},
	"cluster_pvc_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PVCCount()
	},
	"cluster_daemonset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DaemonSetCount()
	},
	"cluster_deployment_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).DeploymentCount()
	},
	"cluster_endpoint_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).EndpointCount()
	},
	"cluster_hpa_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).HPACount()
	},
	"cluster_job_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).JobCount()
	},
	"cluster_statefulset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).StatefulSetCount()
	},
	"cluster_replicaset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).ReplicaSetCount()
	},
	"cluster_service_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).ServiceCount()
	},
	"cluster_secret_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).SecretCount()
	},
	"cluster_pv_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).PVCount()
	},
	"cluster_ingresses_extensions_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).IngressCount()
	},
	"cluster_load1": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).Load1()
	},
	"cluster_load5": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).Load5()
	},
	"cluster_load15": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).Load15()
	},
	"kubesphere_workspace_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).WorkspaceCount()
	},
	"kubesphere_user_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).UserCount()
	},
	"kubesphere_cluser_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Cluster).ClusterCount()
	},
}
