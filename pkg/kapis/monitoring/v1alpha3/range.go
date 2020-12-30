package v1alpha3

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring"
	"strconv"
	"time"
)

type timeRange struct {
	Time  time.Time
	Range monitoring.Range

	isRange bool
}

func (tr timeRange) IsRange() bool {
	return tr.isRange
}

func newTimeRange(ts, start, end, step string) (*timeRange, error) {
	var tr timeRange
	var err error
	var timestamp int64

	// range queryMetrics
	if start != "" && end != "" {
		timestamp, err = strconv.ParseInt(start, 10, 64)
		if err != nil {
			return nil, err
		}
		tr.Range.Start = time.Unix(timestamp, 0)

		timestamp, err = strconv.ParseInt(end, 10, 64)
		if err != nil {
			return nil, err
		}
		tr.Range.End = time.Unix(timestamp, 0)

		if tr.Range.Start.After(tr.Range.End) {
			return nil, fmt.Errorf("invalid time range")
		}

		tr.Range.Step = 10 * time.Minute
		if step != "" {
			if tr.Range.Step, err = time.ParseDuration(step); err != nil {
				return nil, err
			}
		}

		tr.isRange = true
		return &tr, nil
	}

	// instant queryMetrics
	tr.Time = time.Now()
	if ts != "" {
		timestamp, err = strconv.ParseInt(ts, 10, 64)
		if err != nil {
			return nil, err
		}
		tr.Time = time.Unix(timestamp, 0)
	}

	tr.isRange = false
	return &tr, nil
}

func parseTimeRange(request *restful.Request) (*timeRange, error) {
	ts := request.QueryParameter("time")
	start := request.QueryParameter("start")
	end := request.QueryParameter("end")
	step := request.QueryParameter("step")

	return newTimeRange(ts, start, end, step)
}
