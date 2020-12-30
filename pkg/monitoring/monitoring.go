package monitoring

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/prometheus"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
	"time"
)

// Interface get all monitor metrics
type Interface interface {
	Cluster() types.Cluster
	Node(name, filter string) types.Node
	Workspace(name, filter string) types.Workspace
	Namespace(workspace, name, filter string) types.Namespace
	Workload(namespace, kind, filter string) types.Workload
	Pod(node, namespace, workloadKind, workloadName, name, filter string) types.Pod
	Container(namespace, pod, name, filter string) types.Container
	PVC(namespace, storageClass, name, filter string) types.PVC
	Component() types.Component
	Query(expr string) types.MetricQuery
	GetMetadata(namespace string) []Metadata
	GetMetricLabelSet(expr string, start, end time.Time) []map[string]string
}

func New(options *Options) (Interface, error) {
	// TODO: to New different backend by options
	return prometheus.New(&prometheus.Options{Endpoint: options.Endpoint})
}
