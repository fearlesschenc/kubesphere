package prometheus

import (
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type container struct {
	selector string
	api      prometheusapiv1.API
}

func (cli *Client) Container(namespace, pod, name, filter string) types.Container {
	c := &container{api: cli.API}

	if name != "" {
		c.selector = fmt.Sprintf(`pod="%s", namespace="%s", container="%s"`, pod, namespace, name)
	} else {
		c.selector = fmt.Sprintf(`pod="%s", namespace="%s", container=~"%s"`, pod, namespace, filter)
	}

	return c
}

func (c *container) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: c.api}
}

func (c *container) CPUUsage() types.MetricQuery {
	return c.newQuery(fmt.Sprintf(`round(sum by (namespace, pod, container) (irate(container_cpu_usage_seconds_total{job="kubelet", container!="POD", container!="", image!="", %s}[5m])), 0.001)`, c.selector))
}

func (c *container) MemoryUsage() types.MetricQuery {
	return c.newQuery(fmt.Sprintf(`sum by (namespace, pod, container) (container_memory_usage_bytes{job="kubelet", container!="POD", container!="", image!="", %s})`, c.selector))
}

func (c *container) MemoryUsageWoCache() types.MetricQuery {
	return c.newQuery(fmt.Sprintf(`sum by (namespace, pod, container) (container_memory_working_set_bytes{job="kubelet", container!="POD", container!="", image!="", %s})`, c.selector))
}
