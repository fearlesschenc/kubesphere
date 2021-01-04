package prometheus

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type cluster struct {
	api prometheusapiv1.API
}

func (cli *Client) Cluster() types.Cluster {
	return &cluster{api: cli.API}
}

func (c *cluster) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: c.api}
}

func (c *cluster) CPUUtilization() types.MetricQuery {
	return c.newQuery(":node_cpu_utilisation:avg1m")
}

func (c *cluster) CPUUsage() types.MetricQuery {
	return c.newQuery("round(:node_cpu_utilisation:avg1m * sum(node:node_num_cpu:sum), 0.001)")
}

func (c *cluster) CPUTotal() types.MetricQuery {
	return c.newQuery("sum(node:node_num_cpu:sum)")
}

func (c *cluster) MemoryUtilization() types.MetricQuery {
	return c.newQuery(":node_memory_utilisation:")
}

func (c *cluster) MemoryAvailable() types.MetricQuery {
	return c.newQuery("sum(node:node_memory_bytes_available:sum)")
}

func (c *cluster) MemoryTotal() types.MetricQuery {
	return c.newQuery("sum(node:node_memory_bytes_total:sum)")
}

func (c *cluster) MemoryUsageWoCache() types.MetricQuery {
	return c.newQuery("sum(node:node_memory_bytes_total:sum) - sum(node:node_memory_bytes_available:sum)")
}

func (c *cluster) NetUtilization() types.MetricQuery {
	return c.newQuery(":node_net_utilisation:sum_irate")
}

func (c *cluster) NetBytesTransmitted() types.MetricQuery {
	return c.newQuery("sum(node:node_net_bytes_transmitted:sum_irate)")
}

func (c *cluster) NetBytesReceived() types.MetricQuery {
	return c.newQuery("sum(node:node_net_bytes_received:sum_irate)")
}

func (c *cluster) DiskReadIOPS() types.MetricQuery {
	return c.newQuery("sum(node:data_volume_iops_reads:sum)")
}

func (c *cluster) DiskReadThroughput() types.MetricQuery {
	return c.newQuery("sum(node:data_volume_throughput_bytes_read:sum)")
}

func (c *cluster) DiskWriteIOPS() types.MetricQuery {
	return c.newQuery("sum(node:data_volume_iops_writes:sum)")
}

func (c *cluster) DiskWriteThroughput() types.MetricQuery {
	return c.newQuery("sum(node:data_volume_throughput_bytes_written:sum)")
}

func (c *cluster) DiskSizeUsage() types.MetricQuery {
	return c.newQuery(`sum(max(node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"} - node_filesystem_avail_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) by (device, instance))`)
}

func (c *cluster) DiskSizeUtilization() types.MetricQuery {
	return c.newQuery("cluster:disk_utilization:ratio")
}

func (c *cluster) DiskSizeCapacity() types.MetricQuery {
	return c.newQuery(`sum(max(node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) by (device, instance))`)
}

func (c *cluster) DiskSizeAvailable() types.MetricQuery {
	return c.newQuery(`sum(max(node_filesystem_avail_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) by (device, instance))`)
}

func (c *cluster) DiskInodeTotal() types.MetricQuery {
	return c.newQuery("sum(node:node_inodes_total:)")
}

func (c *cluster) DiskInodeUsage() types.MetricQuery {
	return c.newQuery("sum(node:node_inodes_total:) - sum(node:node_inodes_free:)")
}

func (c *cluster) DiskInodeUtilization() types.MetricQuery {
	return c.newQuery("cluster:disk_inode_utilization:ratio")
}

func (c *cluster) PodCount() types.MetricQuery {
	return c.newQuery("cluster:pod:sum")
}

func (c *cluster) PodQuota() types.MetricQuery {
	return c.newQuery(`sum(max(kube_node_status_capacity{resource="pods"}) by (node) unless on (node) (kube_node_status_condition{condition="Ready",status=~"unknown|false"} > 0))`)
}

func (c *cluster) PodUtilization() types.MetricQuery {
	return c.newQuery("cluster:pod_utilization:ratio")
}

func (c *cluster) PodRunning() types.MetricQuery {
	return c.newQuery("cluster:pod_running:count")
}

func (c *cluster) PodSucceed() types.MetricQuery {
	return c.newQuery(`count(kube_pod_info unless on (pod) (kube_pod_status_phase{phase=~"Failed|Pending|Unknown|Running"} > 0) unless on (node) (kube_node_status_condition{condition="Ready",status=~"unknown|false"} > 0))`)
}

func (c *cluster) PodAbnormal() types.MetricQuery {
	return c.newQuery("cluster:pod_abnormal:sum")
}

func (c *cluster) PodAbnormalRatio() types.MetricQuery {
	return c.newQuery(`cluster:pod_abnormal:ratio`)
}

func (c *cluster) NodeTotal() types.MetricQuery {
	return c.newQuery(`sum(kube_node_status_condition{condition="Ready"})`)
}

func (c *cluster) NodeOnline() types.MetricQuery {
	return c.newQuery(`sum(kube_node_status_condition{condition="Ready",status="true"})`)
}

func (c *cluster) NodeOffline() types.MetricQuery {
	return c.newQuery("cluster:node_offline:sum")
}

func (c *cluster) NodeOfflineRatio() types.MetricQuery {
	return c.newQuery(`cluster:node_offline:ratio`)
}

func (c *cluster) NamespaceCount() types.MetricQuery {
	return c.newQuery(`count(kube_namespace_labels)`)
}

func (c *cluster) CronJobCount() types.MetricQuery {
	return c.newQuery(`sum(kube_cronjob_labels)`)
}

func (c *cluster) PVCCount() types.MetricQuery {
	return c.newQuery(`sum(kube_persistentvolumeclaim_info)`)
}

func (c *cluster) DaemonSetCount() types.MetricQuery {
	return c.newQuery(`sum(kube_daemonset_labels)`)
}

func (c *cluster) DeploymentCount() types.MetricQuery {
	return c.newQuery(`sum(kube_deployment_labels)`)
}

func (c *cluster) EndpointCount() types.MetricQuery {
	return c.newQuery(`sum(kube_endpoint_labels)`)
}

func (c *cluster) HPACount() types.MetricQuery {
	return c.newQuery(`sum(kube_hpa_labels)`)
}

func (c *cluster) JobCount() types.MetricQuery {
	return c.newQuery(`sum(kube_job_labels)`)
}

func (c *cluster) StatefulSetCount() types.MetricQuery {
	return c.newQuery(`sum(kube_statefulset_labels)`)
}

func (c *cluster) ReplicaSetCount() types.MetricQuery {
	return c.newQuery(`count(kube_replicaset_labels)`)
}

func (c *cluster) ServiceCount() types.MetricQuery {
	return c.newQuery(`sum(kube_service_info)`)
}

func (c *cluster) SecretCount() types.MetricQuery {
	return c.newQuery(`sum(kube_secret_info)`)
}

func (c *cluster) PVCount() types.MetricQuery {
	return c.newQuery(`sum(kube_persistentvolume_labels)`)
}

func (c *cluster) IngressCount() types.MetricQuery {
	return c.newQuery(`sum(kube_ingress_labels)`)
}

func (c *cluster) Load1() types.MetricQuery {
	return c.newQuery(`sum(node_load1{job="node-exporter"}) / sum(node:node_num_cpu:sum)`)
}

func (c *cluster) Load5() types.MetricQuery {
	return c.newQuery(`sum(node_load5{job="node-exporter"}) / sum(node:node_num_cpu:sum)`)
}

func (c *cluster) Load15() types.MetricQuery {
	return c.newQuery(`sum(node_load15{job="node-exporter"}) / sum(node:node_num_cpu:sum)`)
}

func (c *cluster) WorkspaceCount() types.MetricQuery {
	return c.newQuery(`kubesphere_workspace_count`)
}

func (c *cluster) UserCount() types.MetricQuery {
	return c.newQuery(`kubesphere_user_count`)
}

func (c *cluster) ClusterCount() types.MetricQuery {
	return c.newQuery(`kubesphere_cluster_count`)
}
