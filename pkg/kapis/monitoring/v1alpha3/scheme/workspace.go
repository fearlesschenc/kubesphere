package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var Workspace = MetricScheme{
	"workspace_cpu_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).CPUUsage()
	},
	"workspace_memory_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).MemoryUsage()
	},
	"workspace_net_bytes_transmitted": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).NetBytesTransmitted()
	},
	"workspace_net_bytes_received": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).NetBytesReceived()
	},
	"workspace_pod_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).PodCount()
	},
	"workspace_pod_running_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).PodRunning()
	},
	"workspace_pod_succeeded_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).PodSucceed()
	},
	"workspace_pod_abnormal_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).PodAbnormal()
	},
	"workspace_pod_abnormal_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).PodAbnormalRatio()
	},
	"workspace_ingresses_extensions_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).IngressCount()
	},
	"workspace_cronjob_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).CronJobCount()
	},
	"workspace_pvc_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).PVCCount()
	},
	"workspace_daemonset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).DaemonSetCount()
	},
	"workspace_deployment_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).DeploymentCount()
	},
	"workspace_endpoint_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).EndpointCount()
	},
	"workspace_hpa_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).HPACount()
	},
	"workspace_job_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).JobCount()
	},
	"workspace_statefulset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).StatefulSetCount()
	},
	"workspace_replicaset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).ReplicaSetCount()
	},
	"workspace_service_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).ServiceCount()
	},
	"workspace_secret_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Workspace).SecretCount()
	},
}
