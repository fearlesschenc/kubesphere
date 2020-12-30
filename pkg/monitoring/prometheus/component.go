package prometheus

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type component struct {
	api prometheusapiv1.API
}

func (cli *Client) Component() types.Component {
	return &component{api: cli.API}
}

func (c *component) Etcd() types.Etcd {
	return &etcd{api: c.api}
}

func (c *component) APIServer() types.APIServer {
	return &apiserver{api: c.api}
}

func (c *component) Scheduler() types.Scheduler {
	return &scheduler{api: c.api}
}

type etcd struct {
	api prometheusapiv1.API
}

func (e *etcd) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: e.api}
}

func (e *etcd) ServerList() types.MetricQuery {
	return e.newQuery(`label_replace(up{job="etcd"}, "node_ip", "$1", "instance", "(.*):.*")`)
}

func (e *etcd) ServerTotal() types.MetricQuery {
	return e.newQuery(`count(up{job="etcd"})`)
}

func (e *etcd) ServerUpTotal() types.MetricQuery {
	return e.newQuery(`etcd:up:sum`)
}

func (e *etcd) ServerHasLeader() types.MetricQuery {
	return e.newQuery(`label_replace(etcd_server_has_leader, "node_ip", "$1", "instance", "(.*):.*")`)
}

func (e *etcd) ServerLeaderChanges() types.MetricQuery {
	return e.newQuery(`label_replace(etcd:etcd_server_leader_changes_seen:sum_changes, "node_ip", "$1", "node", "(.*)")`)
}

func (e *etcd) ServerProposalsFailedRate() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_server_proposals_failed:sum_irate)`)
}

func (e *etcd) ServerProposalsAppliedRate() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_server_proposals_applied:sum_irate)`)
}

func (e *etcd) ServerProposalsCommittedRate() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_server_proposals_committed:sum_irate)`)
}

func (e *etcd) ServerProposalsPendingCount() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_server_proposals_pending:sum)`)
}

func (e *etcd) MvccDBSize() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_debugging_mvcc_db_total_size:sum)`)
}

func (e *etcd) NetworkClientGRPCReceivedBytes() types.MetricQuery {
	return e.newQuery(`sum(etcd:etcd_network_client_grpc_received_bytes:sum_irate)`)
}

func (e *etcd) NetworkClientGRPCSentBytes() types.MetricQuery {
	return e.newQuery(`sum(etcd:etcd_network_client_grpc_sent_bytes:sum_irate)`)
}

func (e *etcd) GRPCCallRate() types.MetricQuery {
	return e.newQuery(`sum(etcd:grpc_server_started:sum_irate)`)
}

func (e *etcd) GRPCCallFailedRate() types.MetricQuery {
	return e.newQuery(`sum(etcd:grpc_server_handled:sum_irate)`)
}

func (e *etcd) GRPCServerMSGReceivedRate() types.MetricQuery {
	return e.newQuery(`sum(etcd:grpc_server_msg_received:sum_irate)`)
}

func (e *etcd) GRPCServerMSGSentRate() types.MetricQuery {
	return e.newQuery(`sum(etcd:grpc_server_msg_sent:sum_irate)`)
}

func (e *etcd) DiskWALFsyncDuration() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_disk_wal_fsync_duration:avg)`)
}

func (e *etcd) DiskWALFsyncDurationQuantile() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_disk_wal_fsync_duration:histogram_quantile) by (quantile)`)
}

func (e *etcd) DiskBackendCommitDuration() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_disk_backend_commit_duration:avg)`)
}

func (e *etcd) DiskBackendCommitDurationQuantile() types.MetricQuery {
	return e.newQuery(`avg(etcd:etcd_disk_backend_commit_duration:histogram_quantile) by (quantile)`)
}

type apiserver struct {
	api prometheusapiv1.API
}

func (s *apiserver) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: s.api}
}

func (s *apiserver) Up() types.MetricQuery {
	return s.newQuery(`apiserver:up:sum`)
}

func (s *apiserver) RequestRate() types.MetricQuery {
	return s.newQuery(`apiserver:apiserver_request_total:sum_irate`)
}

func (s *apiserver) RequestByVerbRate() types.MetricQuery {
	return s.newQuery(`apiserver:apiserver_request_total:sum_verb_irate`)
}

func (s *apiserver) RequestLatencies() types.MetricQuery {
	return s.newQuery(`apiserver:apiserver_request_duration:avg`)
}

func (s *apiserver) RequestByVerbLatencies() types.MetricQuery {
	return s.newQuery(`apiserver:apiserver_request_duration:avg_by_verb`)
}

type scheduler struct {
	api prometheusapiv1.API
}

func (s *scheduler) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: s.api}
}

func (s *scheduler) Up() types.MetricQuery {
	return s.newQuery(`scheduler:up:sum`)
}

func (s *scheduler) ScheduleAttempts() types.MetricQuery {
	return s.newQuery(`scheduler:scheduler_schedule_attempts:sum`)
}

func (s *scheduler) ScheduleAttemptRate() types.MetricQuery {
	return s.newQuery(`scheduler:scheduler_schedule_attempts:sum_rate`)
}

func (s *scheduler) E2ESchedulingLatency() types.MetricQuery {
	return s.newQuery(`scheduler:scheduler_e2e_scheduling_duration:avg`)
}

func (s *scheduler) E2ESchedulingLatencyQuantile() types.MetricQuery {
	return s.newQuery(`scheduler:scheduler_e2e_scheduling_duration:histogram_quantile`)
}
