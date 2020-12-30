package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var NamespaceScheme = MetricScheme{
	"namespace_cpu_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).CPUUsage()
	},
	"namespace_memory_usage": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).MemoryUsage()
	},
	"namespace_net_bytes_transmitted": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).NetBytesTransmitted()
	},
	"namespace_net_bytes_received": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).NetBytesReceived()
	},
	"namespace_pod_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).PodCount()
	},
	"namespace_pod_count_hard": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).PodCountHard()
	},
	"namespace_pod_running_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).PodRunning()
	},
	"namespace_pod_succeeded_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).PodSucceed()
	},
	"namespace_pod_abnormal_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).PodAbnormal()
	},
	"namespace_pod_abnormal_ratio": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).PodAbnormalRatio()
	},
	"namespace_memory_limit_hard": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).MemoryLimitHard()
	},
	"namespace_cpu_limit_hard": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).CPULimitHard()
	},
	"namespace_cronjob_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).CronJobCount()
	},
	"namespace_pvc_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).PVCCount()
	},
	"namespace_daemonset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).DaemonSetCount()
	},
	"namespace_deployment_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).DeploymentCount()
	},
	"namespace_endpoint_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).EndpointCount()
	},
	"namespace_hpa_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).HPACount()
	},
	"namespace_job_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).JobCount()
	},
	"namespace_statefulset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).StatefulSetCount()
	},
	"namespace_replicaset_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).ReplicasetCount()
	},
	"namespace_service_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).ServiceCount()
	},
	"namespace_secret_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).SecretCount()
	},
	"namespace_configmap_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).ConfigMapCount()
	},
	"namespace_ingresses_extensions_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).IngressCount()
	},
	"namespace_s2ibuilder_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Namespace).S2IBuilderCount()
	},
}
