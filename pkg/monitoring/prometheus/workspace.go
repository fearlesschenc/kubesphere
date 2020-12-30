package prometheus

import (
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type workspace struct {
	selector string
	api      prometheusapiv1.API
}

func (cli *Client) Workspace(name, filter string) types.Workspace {
	if name != "" {
		return &workspace{selector: fmt.Sprintf(`workspace="%s"`, name), api: cli.API}
	}

	return &workspace{selector: fmt.Sprintf(`workspace=~"%s", workspace!=""`, filter), api: cli.API}
}

func (w *workspace) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: w.api}
}

func (w *workspace) CPUUsage() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`round(sum by (workspace) (namespace:container_cpu_usage_seconds_total:sum_rate{namespace!="", %s}), 0.001)`, w.selector))
}

func (w *workspace) MemoryUsage() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (namespace:container_memory_usage_bytes:sum{namespace!="", %s})`, w.selector))
}

// TODO:
//  "workspace_memory_usage_wo_cache":      `sum by (workspace) (namespace:container_memory_usage_bytes_wo_cache:sum{namespace!="", $1})`,

func (w *workspace) NetBytesTransmitted() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (sum by (namespace) (irate(container_network_transmit_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace) group_left(workspace) kube_namespace_labels{%[1]%}) or on(workspace) max by(workspace) (kube_namespace_labels{%[1]s} * 0)`, w.selector))
}

func (w *workspace) NetBytesReceived() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (sum by (namespace) (irate(container_network_receive_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace) group_left(workspace) kube_namespace_labels{%[1]s}) or on(workspace) max by(workspace) (kube_namespace_labels{%[1]s} * 0)`, w.selector))
}

func (w *workspace) PodCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_pod_status_phase{phase!~"Failed|Succeeded", namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%[1]s})) or on(workspace) max by(workspace) (kube_namespace_labels{%[1]s} * 0)`, w.selector))
}

func (w *workspace) PodRunning() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_pod_status_phase{phase="Running", namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%[1]s})) or on(workspace) max by(workspace) (kube_namespace_labels{%[1]s} * 0)`, w.selector))
}

func (w *workspace) PodSucceed() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_pod_status_phase{phase="Succeeded", namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%[1]s})) or on(workspace) max by(workspace) (kube_namespace_labels{%[1]s} * 0)`, w.selector))
}

func (w *workspace) PodAbnormal() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`count by (workspace) ((kube_pod_info{node!=""} unless on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Succeeded"}>0) unless on (pod, namespace) ((kube_pod_status_ready{job="kube-state-metrics", condition="true"}>0) and on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Running"}>0)) unless on (pod, namespace) (kube_pod_container_status_waiting_reason{job="kube-state-metrics", reason="ContainerCreating"}>0)) * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) PodAbnormalRatio() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`count by (workspace) ((kube_pod_info{node!=""} unless on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Succeeded"}>0) unless on (pod, namespace) ((kube_pod_status_ready{job="kube-state-metrics", condition="true"}>0) and on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Running"}>0)) unless on (pod, namespace) (kube_pod_container_status_waiting_reason{job="kube-state-metrics", reason="ContainerCreating"}>0)) * on (namespace) group_left(workspace) kube_namespace_labels{%[1]s}) / sum by (workspace) (kube_pod_status_phase{phase!="Succeeded", namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%[1]s}))`, w.selector))
}

// TODO: add namespace count
func (w *workspace) ApplicationCount() types.MetricQuery {
	return w.newQuery("")
}

func (w *workspace) CronJobCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_cronjob_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) PVCCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_persistentvolumeclaim_info{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) DaemonSetCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_daemonset_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) DeploymentCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_deployment_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) EndpointCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_endpoint_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) HPACount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_hpa_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) JobCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_job_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) StatefulSetCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_statefulset_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) ServiceCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_service_info{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) SecretCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_secret_info{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) ReplicaSetCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`count by (workspace) (kube_replicaset_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}

func (w *workspace) IngressCount() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`sum by (workspace) (kube_ingress_labels{namespace!=""} * on (namespace) group_left(workspace)(kube_namespace_labels{%s}))`, w.selector))
}
