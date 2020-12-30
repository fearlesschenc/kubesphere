package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var SchedulerScheme = MetricScheme{
	"scheduler_up_sum": func(obj types.Object) types.MetricQuery {
		return obj.(types.Scheduler).Up()
	},
	"scheduler_schedule_attempts": func(obj types.Object) types.MetricQuery {
		return obj.(types.Scheduler).ScheduleAttempts()
	},
	"scheduler_schedule_attempt_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.Scheduler).ScheduleAttemptRate()
	},
	"scheduler_e2e_scheduling_latency": func(obj types.Object) types.MetricQuery {
		return obj.(types.Scheduler).E2ESchedulingLatency()
	},
	"scheduler_e2e_scheduling_latency_quantile": func(obj types.Object) types.MetricQuery {
		return obj.(types.Scheduler).E2ESchedulingLatencyQuantile()
	},
}
