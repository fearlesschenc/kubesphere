package prometheus

import (
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type workload struct {
	selector     string
	kindSelector map[string]string
	api          prometheusapiv1.API
}

func (cli *Client) Workload(namespace, kind, filter string) types.Workload {
	wl := &workload{
		selector:     fmt.Sprintf(`namespace="%s", workload=~"%s:(%s)"`, namespace, kind, filter),
		kindSelector: map[string]string{},
		api:          cli.API,
	}

	wl.kindSelector["deployment"] = fmt.Sprintf(`namespace="%s", deployment!="", deployment=~"%s"`, namespace, filter)
	wl.kindSelector["statefulset"] = fmt.Sprintf(`namespace="%s", statefulset!="", statefulset=~"%s"`, namespace, filter)
	wl.kindSelector["daemonset"] = fmt.Sprintf(`namespace="%s", daemonset!="", daemonset=~"%s"`, namespace, filter)

	return wl
}

func (w *workload) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: w.api}
}

func (w *workload) CPUUsage() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`round(namespace:workload_cpu_usage:sum{%s}, 0.001)`, w.selector))
}

func (w *workload) MemoryUsage() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`namespace:workload_memory_usage:sum{%s}`, w.selector))
}

func (w *workload) NetBytesTransmitted() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`namespace:workload_net_bytes_transmitted:sum_irate{%s}`, w.selector))
}

func (w *workload) NetBytesReceived() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`namespace:workload_net_bytes_received:sum_irate{%s}`, w.selector))
}

func (w *workload) DeploymentReplica() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`label_join(sum (label_join(label_replace(kube_deployment_spec_replicas{%s}, "owner_kind", "Deployment", "", ""), "workload", "", "deployment")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`, w.kindSelector["deployment"]))
}

func (w *workload) DeploymentReplicaAvailable() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`label_join(sum (label_join(label_replace(kube_deployment_status_replicas_available{%s}, "owner_kind", "Deployment", "", ""), "workload", "", "deployment")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`, w.kindSelector["deployment"]))
}

func (w *workload) DeploymentUnavailableReplicasRatio() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`namespace:deployment_unavailable_replicas:ratio{%s}`, w.selector))
}

func (w *workload) StatefulSetReplica() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`label_join(sum (label_join(label_replace(kube_statefulset_replicas{%s}, "owner_kind", "StatefulSet", "", ""), "workload", "", "statefulset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`, w.kindSelector["statefulset"]))
}

func (w *workload) StatefulSetReplicaAvailable() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`label_join(sum (label_join(label_replace(kube_statefulset_status_replicas_current{%s}, "owner_kind", "StatefulSet", "", ""), "workload", "", "statefulset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`, w.kindSelector["statefulset"]))
}

func (w *workload) StatefulSetUnavailableReplicaRatio() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`namespace:statefulset_unavailable_replicas:ratio{%s}`, w.selector))
}

func (w *workload) DaemonSetReplica() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`label_join(sum (label_join(label_replace(kube_daemonset_status_desired_number_scheduled{%s}, "owner_kind", "DaemonSet", "", ""), "workload", "", "daemonset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`, w.kindSelector["daemonset"]))
}

func (w *workload) DaemonSetReplicaAvailable() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`label_join(sum (label_join(label_replace(kube_daemonset_status_number_available{%s}, "owner_kind", "DaemonSet", "", ""), "workload", "", "daemonset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`, w.kindSelector["daemonset"]))
}

func (w *workload) DaemonSetUnavailableReplicasRatio() types.MetricQuery {
	return w.newQuery(fmt.Sprintf(`namespace:daemonset_unavailable_replicas:ratio{%s}`, w.selector))
}
