package prometheus

import (
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type namespace struct {
	selector string
	api      prometheusapiv1.API
}

func (cli *Client) Namespace(workspace, name, filter string) types.Namespace {
	if workspace != "" {
		return &namespace{selector: fmt.Sprintf(`workspace="%s", namespace=~"%s"`, workspace, filter)}
	}

	if name != "" {
		return &namespace{selector: fmt.Sprintf(`namespace="%s"`, name), api: cli.API}
	}

	return &namespace{selector: fmt.Sprintf(`namespace=~"%s"`, filter), api: cli.API}
}

func (ns *namespace) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: ns.api}
}

func (ns *namespace) CPUUsage() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`round(namespace:container_cpu_usage_seconds_total:sum_rate{namespace!="", %s}, 0.001)`, ns.selector))
}

func (ns *namespace) CPULimitHard() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`min by (namespace) (kube_resourcequota{resourcequota!="quota", type="hard", namespace!="", resource="limits.cpu"} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) MemoryUsage() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`namespace:container_memory_usage_bytes:sum{namespace!="", %s}`, ns.selector))
}

func (ns *namespace) MemoryLimitHard() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`min by (namespace) (kube_resourcequota{resourcequota!="quota", type="hard", namespace!="", resource="limits.memory"} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) NetBytesTransmitted() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (irate(container_network_transmit_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m]) * on (namespace) group_left(workspace) kube_namespace_labels{%[1]s}) or on(namespace) max by(namespace) (kube_namespace_labels{%[1]s} * 0)`, ns.selector))
}

func (ns *namespace) NetBytesReceived() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (irate(container_network_receive_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m]) * on (namespace) group_left(workspace) kube_namespace_labels{%[1]s}) or on(namespace) max by(namespace) (kube_namespace_labels{%[1]s} * 0)`, ns.selector))
}

func (ns *namespace) PodCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_pod_status_phase{phase!~"Failed|Succeeded", namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%[1]s}) or on(namespace) max by(namespace) (kube_namespace_labels{%[1]s} * 0)`, ns.selector))
}

func (ns *namespace) PodCountHard() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`min by (namespace) (kube_resourcequota{resourcequota!="quota", type="hard", namespace!="", resource="count/pods"} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) PodRunning() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_pod_status_phase{phase="Running", namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%[1]s}) or on(namespace) max by(namespace) (kube_namespace_labels{%[1]s} * 0)`, ns.selector))
}

func (ns *namespace) PodSucceed() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_pod_status_phase{phase="Succeeded", namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%[1]s}) or on(namespace) max by(namespace) (kube_namespace_labels{%[1]s} * 0)`, ns.selector))
}

func (ns *namespace) PodAbnormal() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`namespace:pod_abnormal:count{namespace!="", %s}`, ns.selector))
}

func (ns *namespace) PodAbnormalRatio() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`namespace:pod_abnormal:ratio{namespace!="", %s}`, ns.selector))
}

func (ns *namespace) CronJobCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_cronjob_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) PVCCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_persistentvolumeclaim_info{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) DaemonSetCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_daemonset_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) DeploymentCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_deployment_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) EndpointCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_endpoint_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) HPACount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_hpa_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) JobCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_job_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) ReplicasetCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`count by (namespace) (kube_replicaset_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) StatefulSetCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_statefulset_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) ServiceCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_service_info{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) SecretCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_secret_info{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) ConfigMapCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_configmap_info{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) IngressCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (kube_ingress_labels{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}

func (ns *namespace) S2IBuilderCount() types.MetricQuery {
	return ns.newQuery(fmt.Sprintf(`sum by (namespace) (s2i_s2ibuilder_created{namespace!=""} * on (namespace) group_left(workspace) kube_namespace_labels{%s})`, ns.selector))
}
