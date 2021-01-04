package prometheus

import (
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type node struct {
	selector string
	api      prometheusapiv1.API
}

func (cli *Client) Node(name, filter string) types.Node {
	n := &node{api: cli.API}

	if name != "" {
		n.selector = fmt.Sprintf(`node="%s"`, name)
	} else {
		n.selector = fmt.Sprintf(`node=~"%s"`, filter)
	}

	return n
}

func (n *node) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: n.api}
}

func (n *node) CPUUtilization() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_cpu_utilisation:avg1m{%s}", n.selector))
}

func (n *node) CPUUsage() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf(`round(node:node_cpu_utilisation:avg1m{%[1]s} * node:node_num_cpu:sum{%[1]s}, 0.001)`, n.selector))
}

func (n *node) CPUTotal() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_num_cpu:sum{%s}", n.selector))
}

func (n *node) MemoryUtilization() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_memory_utilisation:{%s}", n.selector))
}

func (n *node) MemoryAvailable() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_memory_bytes_available:sum{%s}", n.selector))
}

func (n *node) MemoryTotal() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_memory_bytes_total:sum{%s}", n.selector))
}

func (n *node) MemoryUsageWoCache() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_memory_bytes_total:sum{%[1]s} - node:node_memory_bytes_available:sum{%[1]s}", n.selector))
}

func (n *node) NetUtilization() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_net_utilisation:sum_irate{%s}", n.selector))
}

func (n *node) NetBytesTransmitted() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_net_bytes_transmitted:sum_irate{%s}", n.selector))
}

func (n *node) NetBytesReceived() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_net_bytes_received:sum_irate{%s}", n.selector))
}

func (n *node) DiskReadIOPS() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:data_volume_iops_reads:sum{%s}", n.selector))
}

func (n *node) DiskReadThroughput() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:data_volume_throughput_bytes_read:sum{%s}", n.selector))
}

func (n *node) DiskWriteIOPS() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:data_volume_iops_writes:sum{%s}", n.selector))
}

func (n *node) DiskWriteThroughput() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:data_volume_throughput_bytes_written:sum{%s}", n.selector))
}

func (n *node) DiskSizeUsage() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf(`sum(max((node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"} - node_filesystem_avail_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) * on (namespace, pod) group_left(node) node_namespace_pod:kube_pod_info:{%s}) by (device, node)) by (node)`, n.selector))
}

func (n *node) DiskSizeUtilization() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:disk_space_utilization:ratio{%s}", n.selector))
}

func (n *node) DiskSizeCapacity() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf(`sum(max(node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"} * on (namespace, pod) group_left(node) node_namespace_pod:kube_pod_info:{%s}) by (device, node)) by (node)`, n.selector))
}

func (n *node) DiskSizeAvailable() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:disk_space_available:{%s}", n.selector))
}

func (n *node) DiskInodeTotal() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_inodes_total:{%s}", n.selector))
}

func (n *node) DiskInodeUsage() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:node_inodes_total:{%[1]s} - node:node_inodes_free:{%[1]s}", n.selector))
}

func (n *node) DiskInodeUtilization() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:disk_inode_utilization:ratio{%s}", n.selector))
}

func (n *node) PodCount() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:pod_count:sum{%s}", n.selector))
}

func (n *node) PodQuota() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf(`max(kube_node_status_capacity{resource="pods", %s}) by (node) unless on (node) (kube_node_status_condition{condition="Ready",status=~"unknown|false"} > 0)`, n.selector))
}

func (n *node) PodUtilization() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:pod_utilization:ratio{%s}", n.selector))
}

func (n *node) PodRunning() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:pod_running:count{%s}", n.selector))
}

func (n *node) PodSucceed() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:pod_succeeded:count{%s}", n.selector))
}

func (n *node) PodAbnormal() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:pod_abnormal:count{%s}", n.selector))
}

func (n *node) PodAbnormalRatio() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:pod_abnormal:ratio{%s}", n.selector))
}

func (n *node) Load1() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:load1:ratio{%s}", n.selector))
}

func (n *node) Load5() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:load5:ratio{%s}", n.selector))
}

func (n *node) Load15() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf("node:load15:ratio{%s}", n.selector))
}

func (n *node) PlegQuantile() types.MetricQuery {
	return n.newQuery(
		fmt.Sprintf(`node_quantile:kubelet_pleg_relist_duration_seconds:histogram_quantile{%s}`, n.selector))
}
