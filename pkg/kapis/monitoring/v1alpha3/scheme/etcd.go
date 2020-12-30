package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var EtcdScheme = MetricScheme{
	"etcd_server_list": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerList()
	},
	"etcd_server_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerTotal()
	},
	"etcd_server_up_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerUpTotal()
	},
	"etcd_server_has_leader": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerHasLeader()
	},
	"etcd_server_leader_changes": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerLeaderChanges()
	},
	"etcd_server_proposals_failed_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerProposalsFailedRate()
	},
	"etcd_server_proposals_applied_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerProposalsAppliedRate()
	},
	"etcd_server_proposals_committed_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerProposalsCommittedRate()
	},
	"etcd_server_proposals_pending_count": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).ServerProposalsPendingCount()
	},
	"etcd_mvcc_db_size": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).MvccDBSize()
	},
	"etcd_network_client_grpc_received_bytes": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).NetworkClientGRPCReceivedBytes()
	},
	"etcd_network_client_grpc_sent_bytes": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).NetworkClientGRPCSentBytes()
	},
	"etcd_grpc_call_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).GRPCCallRate()
	},
	"etcd_grpc_call_failed_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).GRPCCallFailedRate()
	},
	"etcd_grpc_server_msg_received_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).GRPCServerMSGReceivedRate()
	},
	"etcd_grpc_server_msg_sent_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).GRPCServerMSGSentRate()
	},
	"etcd_disk_wal_fsync_duration": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).DiskWALFsyncDuration()
	},
	"etcd_disk_wal_fsync_duration_quantile": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).DiskWALFsyncDurationQuantile()
	},
	"etcd_disk_backend_commit_duration": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).DiskBackendCommitDuration()
	},
	"etcd_disk_backend_commit_duration_quantile": func(obj types.Object) types.MetricQuery {
		return obj.(types.Etcd).DiskBackendCommitDurationQuantile()
	},
}
