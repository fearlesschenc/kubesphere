package prometheus

import (
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type pod struct {
	selector         string
	workloadSelector string
	api              prometheusapiv1.API
}

func (cli *Client) Pod(node, namespace, workloadKind, workloadName, name, filter string) types.Pod {
	p := &pod{api: cli.API}

	if workloadName != "" {
		switch workloadKind {
		case "deployment":
			p.workloadSelector = fmt.Sprintf(`owner_kind="ReplicaSet", owner_name=~"^%s-[^-]{1,10}$"`, workloadName)
		case "statefulset":
			p.workloadSelector = fmt.Sprintf(`owner_kind="StatefulSet", owner_name="%s"`, workloadName)
		case "daemonset":
			p.workloadSelector = fmt.Sprintf(`owner_kind="DaemonSet", owner_name="%s"`, workloadName)
		}
	}

	if namespace != "" {
		if name != "" {
			p.selector = fmt.Sprintf(`pod="%s", namespace="%s"`, name, namespace)
		} else {
			p.selector = fmt.Sprintf(`pod=~"%s", namespace="%s"`, filter, namespace)
		}
	} else if node != "" {
		if name != "" {
			p.selector = fmt.Sprintf(`pod="%s", node="%s"`, name, node)
		} else {
			p.selector = fmt.Sprintf(`pod=~"%s", node="%s"`, filter, node)
		}
	}

	return p
}

func (p *pod) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: p.api}
}

func (p *pod) CPUUsage() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`round(sum by (namespace, pod) (irate(container_cpu_usage_seconds_total{job="kubelet", pod!="", image!=""}[5m])) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{%s} * on (namespace, pod) group_left(node) kube_pod_info{%s}, 0.001)`, p.workloadSelector, p.selector))
}

func (p *pod) MemoryUsage() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`sum by (namespace, pod) (container_memory_usage_bytes{job="kubelet", pod!="", image!=""}) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{%s} * on (namespace, pod) group_left(node) kube_pod_info{%s}`, p.workloadSelector, p.selector))
}

func (p *pod) NetBytesTransmitted() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`sum by (namespace, pod) (irate(container_network_transmit_bytes_total{pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{%s} * on (namespace, pod) group_left(node) kube_pod_info{%s}`, p.workloadSelector, p.selector))
}

func (p *pod) NetBytesReceived() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`sum by (namespace, pod) (irate(container_network_receive_bytes_total{pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{%s} * on (namespace, pod) group_left(node) kube_pod_info{%s}`, p.workloadSelector, p.selector))
}
