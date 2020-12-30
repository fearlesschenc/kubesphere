package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var APIServerScheme = MetricScheme{
	"apiserver_up_sum": func(obj types.Object) types.MetricQuery {
		return obj.(types.APIServer).Up()
	},
	"apiserver_request_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.APIServer).RequestRate()
	},
	"apiserver_request_by_verb_rate": func(obj types.Object) types.MetricQuery {
		return obj.(types.APIServer).RequestByVerbRate()
	},
	"apiserver_request_latencies": func(obj types.Object) types.MetricQuery {
		return obj.(types.APIServer).RequestLatencies()
	},
	"apiserver_request_by_verb_latencies": func(obj types.Object) types.MetricQuery {
		return obj.(types.APIServer).RequestByVerbLatencies()
	},
}
