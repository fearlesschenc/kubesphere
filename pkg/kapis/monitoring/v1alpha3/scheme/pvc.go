package scheme

import (
	"github.com/fearlesschenc/kubesphere/pkg/monitoring/types"
)

var PVCScheme = MetricScheme{
	"pvc_inodes_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).InodeAvailable()
	},
	"pvc_inodes_used": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).InodeUsed()
	},
	"pvc_inodes_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).InodeTotal()
	},
	"pvc_inodes_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).InodeUtilization()
	},
	"pvc_bytes_available": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).BytesAvailable()
	},
	"pvc_bytes_used": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).BytesUsed()
	},
	"pvc_bytes_total": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).BytesTotal()
	},
	"pvc_bytes_utilisation": func(obj types.Object) types.MetricQuery {
		return obj.(types.PVC).BytesUtilization()
	},
}
