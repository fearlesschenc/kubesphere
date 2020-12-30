package prometheus

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

func (cli *Client) Query(expr string) types.MetricQuery {
	return &Query{
		query: expr,
		api:   cli.API,
	}
}
