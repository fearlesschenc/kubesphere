package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var WorkloadScheme = MetricScheme{
	"workload_cpu_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).CPUUsage()
	},
	"workload_memory_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).MemoryUsage()
	},
	"workload_net_bytes_transmitted": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).NetBytesTransmitted()
	},
	"workload_net_bytes_received": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).NetBytesReceived()
	},
	"workload_deployment_replica": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).DeploymentReplica()
	},
	"workload_deployment_replica_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).DeploymentReplicaAvailable()
	},
	"workload_deployment_unavailable_replicas_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).DeploymentUnavailableReplicasRatio()
	},
	"workload_statefulset_replica": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).StatefulSetReplica()
	},
	"workload_statefulset_replica_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).StatefulSetReplicaAvailable()
	},
	"workload_statefulset_unavailable_replicas_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).StatefulSetUnavailableReplicaRatio()
	},
	"workload_daemonset_replica": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).DaemonSetReplica()
	},
	"workload_daemonset_replica_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).DaemonSetReplicaAvailable()
	},
	"workload_daemonset_unavailable_replicas_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workload).DaemonSetUnavailableReplicasRatio()
	},
}
