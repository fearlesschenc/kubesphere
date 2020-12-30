package prometheus

import (
	"context"
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapi "github.com/prometheus/client_golang/api"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"k8s.io/klog"
	"time"
)

type Options struct {
	Endpoint string `json:"endpoint"`
}

func (s *Options) Validate() error {
	return nil
}

type Client struct {
	prometheusapiv1.API
}

func New(options *Options) (*Client, error) {
	client, err := prometheusapi.NewClient(prometheusapi.Config{Address: options.Endpoint})
	if err != nil {
		return nil, err
	}

	// TODO: prefetch
	return &Client{API: prometheusapiv1.NewAPI(client)}, nil
}

// TODO: this interface bound to implementation
func (cli *Client) GetMetadata(namespace string) []types.Metadata {
	var meta []types.Metadata

	// Filter metrics available to members of this namespace
	matchTarget := fmt.Sprintf("{namespace=\"%s\"}", namespace)
	items, err := cli.TargetsMetadata(context.Background(), matchTarget, "", "")
	if err != nil {
		klog.Error(err)
		return meta
	}

	// Deduplication
	set := make(map[string]bool)
	for _, item := range items {
		_, ok := set[item.Metric]
		if !ok {
			set[item.Metric] = true
			meta = append(meta, types.Metadata{
				MetricName: item.Metric,
				Type:       string(item.Type),
				Help:       item.Help,
			})
		}
	}

	return meta
}

func (cli *Client) GetMetricLabelSet(expr string, start, end time.Time) []map[string]string {
	var res []map[string]string

	labelSet, _, err := cli.Series(context.Background(), []string{expr}, start, end)
	if err != nil {
		klog.Error(err)
		return []map[string]string{}
	}

	for _, item := range labelSet {
		var tmp = map[string]string{}
		for key, val := range item {
			if key == "__name__" {
				continue
			}
			tmp[string(key)] = string(val)
		}

		res = append(res, tmp)
	}

	return res
}
