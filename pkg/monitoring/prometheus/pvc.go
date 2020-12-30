package prometheus

import (
	"fmt"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	prometheusapiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type pvc struct {
	selector string
	api      prometheusapiv1.API
}

func (cli *Client) PVC(namespace, storageClass, name, filter string) types.PVC {
	p := &pvc{api: cli.API}

	if namespace != "" {
		if name != "" {
			p.selector = fmt.Sprintf(`namespace="%s", persistentvolumeclaim="%s"`, namespace, name)
		} else {
			p.selector = fmt.Sprintf(`namespace="%s", persistentvolumeclaim=~"%s"`, namespace, filter)
		}

		return p
	}

	if storageClass != "" {
		p.selector = fmt.Sprintf(`storageclass="%s", persistentvolumeclaim=~"%s"`, storageClass, filter)
	}

	return p
}

func (p *pvc) newQuery(query string) types.MetricQuery {
	return &Query{query: query, api: p.api}
}

func (p *pvc) InodeAvailable() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes_free) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}

func (p *pvc) InodeUsed() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes_used) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}

func (p *pvc) InodeTotal() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}

func (p *pvc) InodeUtilization() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes_used / kubelet_volume_stats_inodes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}

func (p *pvc) BytesAvailable() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_available_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}

func (p *pvc) BytesUsed() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_used_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}

func (p *pvc) BytesTotal() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_capacity_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}

func (p *pvc) BytesUtilization() types.MetricQuery {
	return p.newQuery(fmt.Sprintf(`max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_used_bytes / kubelet_volume_stats_capacity_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{%s}`, p.selector))
}
