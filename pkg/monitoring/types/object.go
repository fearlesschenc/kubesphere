package types

type Object interface{}

type Cluster interface {
	Object

	// cpu
	CPUUtilization() MetricQuery
	CPUUsage() MetricQuery
	CPUTotal() MetricQuery

	// memory
	MemoryUtilization() MetricQuery
	MemoryAvailable() MetricQuery
	MemoryTotal() MetricQuery
	MemoryUsageWoCache() MetricQuery

	// net
	NetUtilization() MetricQuery
	NetBytesTransmitted() MetricQuery
	NetBytesReceived() MetricQuery

	// disk
	DiskReadIOPS() MetricQuery
	DiskReadThroughput() MetricQuery
	DiskWriteIOPS() MetricQuery
	DiskWriteThroughput() MetricQuery
	DiskSizeUsage() MetricQuery
	DiskSizeUtilization() MetricQuery
	DiskSizeCapacity() MetricQuery
	DiskSizeAvailable() MetricQuery
	DiskInodeTotal() MetricQuery
	DiskInodeUsage() MetricQuery
	DiskInodeUtilization() MetricQuery

	// pod
	PodCount() MetricQuery
	PodQuota() MetricQuery
	PodUtilization() MetricQuery
	PodRunning() MetricQuery
	PodSucceed() MetricQuery
	PodAbnormal() MetricQuery
	PodAbnormalRatio() MetricQuery

	// node
	NodeTotal() MetricQuery
	NodeOnline() MetricQuery
	NodeOffline() MetricQuery
	NodeOfflineRatio() MetricQuery

	// kubernetes object count
	NamespaceCount() MetricQuery
	CronJobCount() MetricQuery
	PVCCount() MetricQuery
	DaemonSetCount() MetricQuery
	DeploymentCount() MetricQuery
	EndpointCount() MetricQuery
	HPACount() MetricQuery
	JobCount() MetricQuery
	StatefulSetCount() MetricQuery
	ReplicaSetCount() MetricQuery
	ServiceCount() MetricQuery
	SecretCount() MetricQuery
	PVCount() MetricQuery
	IngressCount() MetricQuery

	// load
	Load1() MetricQuery
	Load5() MetricQuery
	Load15() MetricQuery

	WorkspaceCount() MetricQuery
	UserCount() MetricQuery
	ClusterCount() MetricQuery
}

type Node interface {
	Object
	// cpu
	CPUUtilization() MetricQuery
	CPUUsage() MetricQuery
	CPUTotal() MetricQuery

	// memory
	MemoryUtilization() MetricQuery
	MemoryAvailable() MetricQuery
	MemoryTotal() MetricQuery
	MemoryUsageWoCache() MetricQuery

	// net
	NetUtilization() MetricQuery
	NetBytesTransmitted() MetricQuery
	NetBytesReceived() MetricQuery

	// disk
	DiskReadIOPS() MetricQuery
	DiskReadThroughput() MetricQuery
	DiskWriteIOPS() MetricQuery
	DiskWriteThroughput() MetricQuery
	DiskSizeUsage() MetricQuery
	DiskSizeUtilization() MetricQuery
	DiskSizeCapacity() MetricQuery
	DiskSizeAvailable() MetricQuery
	DiskInodeTotal() MetricQuery
	DiskInodeUsage() MetricQuery
	DiskInodeUtilization() MetricQuery

	// pod
	PodCount() MetricQuery
	PodQuota() MetricQuery
	PodUtilization() MetricQuery
	PodRunning() MetricQuery
	PodSucceed() MetricQuery
	PodAbnormal() MetricQuery
	PodAbnormalRatio() MetricQuery

	// load
	Load1() MetricQuery
	Load5() MetricQuery
	Load15() MetricQuery

	PlegQuantile() MetricQuery
}

type Workspace interface {
	Object
	// cpu
	CPUUsage() MetricQuery

	// memory
	MemoryUsage() MetricQuery
	MemoryUsageWoCache() MetricQuery

	// net
	NetBytesTransmitted() MetricQuery
	NetBytesReceived() MetricQuery

	// pod
	PodCount() MetricQuery
	PodRunning() MetricQuery
	PodSucceed() MetricQuery
	PodAbnormal() MetricQuery
	PodAbnormalRatio() MetricQuery

	// kubernetes object count
	NamespaceCount() MetricQuery
	DevopsCount() MetricQuery
	MemberCount() MetricQuery
	RoleCount() MetricQuery
	CronJobCount() MetricQuery
	PVCCount() MetricQuery
	DaemonSetCount() MetricQuery
	DeploymentCount() MetricQuery
	EndpointCount() MetricQuery
	HPACount() MetricQuery
	JobCount() MetricQuery
	StatefulSetCount() MetricQuery
	ReplicaSetCount() MetricQuery
	ServiceCount() MetricQuery
	SecretCount() MetricQuery
	IngressCount() MetricQuery
}

type Namespace interface {
	Object
	// cpu
	CPUUsage() MetricQuery
	CPULimitHard() MetricQuery

	// memory
	MemoryUsage() MetricQuery
	MemoryUsageWoCache() MetricQuery
	MemoryLimitHard() MetricQuery

	// net
	NetBytesTransmitted() MetricQuery
	NetBytesReceived() MetricQuery

	// pod
	PodCount() MetricQuery
	PodCountHard() MetricQuery
	PodRunning() MetricQuery
	PodSucceed() MetricQuery
	PodAbnormal() MetricQuery
	PodAbnormalRatio() MetricQuery

	// kubernetes object count
	CronJobCount() MetricQuery
	PVCCount() MetricQuery
	DaemonSetCount() MetricQuery
	DeploymentCount() MetricQuery
	EndpointCount() MetricQuery
	HPACount() MetricQuery
	JobCount() MetricQuery
	ReplicasetCount() MetricQuery
	StatefulSetCount() MetricQuery
	ServiceCount() MetricQuery
	SecretCount() MetricQuery
	ConfigMapCount() MetricQuery
	IngressCount() MetricQuery
	S2IBuilderCount() MetricQuery
}

type Workload interface {
	Object

	CPUUsage() MetricQuery
	MemoryUsage() MetricQuery
	MemoryUsageWoCache() MetricQuery
	NetBytesTransmitted() MetricQuery
	NetBytesReceived() MetricQuery

	DeploymentReplica() MetricQuery
	DeploymentReplicaAvailable() MetricQuery
	DeploymentUnavailableReplicasRatio() MetricQuery

	StatefulSetReplica() MetricQuery
	StatefulSetReplicaAvailable() MetricQuery
	StatefulSetUnavailableReplicaRatio() MetricQuery

	DaemonSetReplica() MetricQuery
	DaemonSetReplicaAvailable() MetricQuery
	DaemonSetUnavailableReplicasRatio() MetricQuery
}

type Pod interface {
	Object

	CPUUsage() MetricQuery
	MemoryUsage() MetricQuery
	MemoryUsageWoCache() MetricQuery
	NetBytesTransmitted() MetricQuery
	NetBytesReceived() MetricQuery
}

type Container interface {
	Object

	CPUUsage() MetricQuery
	MemoryUsage() MetricQuery
	MemoryUsageWoCache() MetricQuery
}

type PVC interface {
	Object

	InodeAvailable() MetricQuery
	InodeUsed() MetricQuery
	InodeTotal() MetricQuery
	InodeUtilization() MetricQuery
	BytesAvailable() MetricQuery
	BytesUsed() MetricQuery
	BytesTotal() MetricQuery
	BytesUtilization() MetricQuery
}

type Component interface {
	Etcd() Etcd
	APIServer() APIServer
	Scheduler() Scheduler
}

type Etcd interface {
	Object

	// server
	ServerList() MetricQuery
	ServerTotal() MetricQuery
	ServerUpTotal() MetricQuery
	ServerHasLeader() MetricQuery
	ServerLeaderChanges() MetricQuery
	ServerProposalsFailedRate() MetricQuery
	ServerProposalsAppliedRate() MetricQuery
	ServerProposalsCommittedRate() MetricQuery
	ServerProposalsPendingCount() MetricQuery

	MvccDBSize() MetricQuery
	NetworkClientGRPCReceivedBytes() MetricQuery
	NetworkClientGRPCSentBytes() MetricQuery
	GRPCCallRate() MetricQuery
	GRPCCallFailedRate() MetricQuery
	GRPCServerMSGReceivedRate() MetricQuery
	GRPCServerMSGSentRate() MetricQuery
	DiskWALFsyncDuration() MetricQuery
	DiskWALFsyncDurationQuantile() MetricQuery
	DiskBackendCommitDuration() MetricQuery
	DiskBackendCommitDurationQuantile() MetricQuery
}

type APIServer interface {
	Object

	Up() MetricQuery
	RequestRate() MetricQuery
	RequestByVerbRate() MetricQuery
	RequestLatencies() MetricQuery
	RequestByVerbLatencies() MetricQuery
}

type Scheduler interface {
	Object

	Up() MetricQuery
	ScheduleAttempts() MetricQuery
	ScheduleAttemptRate() MetricQuery
	E2ESchedulingLatency() MetricQuery
	E2ESchedulingLatencyQuantile() MetricQuery
}
