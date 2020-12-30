package v1alpha3

import (
	"github.com/emicklei/go-restful"
	"github.com/fearlesschenc/kubesphere/pkg/monitoring"
	"math"
	"sort"
	"strconv"
)

type wrapper struct {
	monitoring.MetricData
	identifier, order string
}

func (w wrapper) Len() int {
	return len(w.MetricValues)
}

func (w wrapper) Less(i, j int) bool {
	p := w.MetricValues[i]
	q := w.MetricValues[j]

	// Place Nil to the tail.
	if p.Sample == nil && q.Sample != nil {
		return false
	}
	if p.Sample != nil && q.Sample == nil {
		return true
	}

	// If both Samples are Nil or have the same metric value, sort by resource name
	if p.Sample == q.Sample || p.Sample[1] == q.Sample[1] {
		return p.Metadata[w.identifier] < q.Metadata[w.identifier]
	}
	// Place NaN to the tail (NaN takes precedence over Nil).
	if math.IsNaN(p.Sample[1]) != math.IsNaN(q.Sample[1]) {
		return !math.IsNaN(p.Sample[1])
	}

	switch w.order {
	case OrderAscending:
		return p.Sample.Value() < q.Sample.Value()
	default:
		return p.Sample.Value() > q.Sample.Value()
	}
}

func (w wrapper) Swap(i, j int) {
	w.MetricValues[i], w.MetricValues[j] = w.MetricValues[j], w.MetricValues[i]
}

// SortMetrics sorts a group of resources by a given metric. Range queryMetrics doesn't support ranking.
// Example:
//
// Before sorting:
// |  ID | Metric 1  |  Metric 2 |  Metric 3 |
// |  a  |     1     |     XL    |           |
// |  b  |     1     |     S     |           |
// |  c  |     3     |     M     |           |
//
// After sorting: target=metric_2, order=asc, identifier=id
// |  ID | Metric 1  |  Metric 2 (asc) |  Metric 3 |
// |  a  |     1     |        XL       |           |
// |  c  |     3     |        M        |           |
// |  b  |     1     |        S        |           |
func (raw *Metrics) Sort(target, order, identifier string) *Metrics {
	if target == "" || identifier == "" || len(raw.Results) == 0 {
		return raw
	}

	resourceSet := make(map[string]bool)    // resource set records possible values of the identifier
	resourceOrdinal := make(map[string]int) // resource-ordinal map

	ordinal := 0
	for _, item := range raw.Results {
		if item.MetricType != monitoring.MetricTypeVector || item.Error != "" {
			continue
		}

		if item.Name == target {
			sort.Sort(wrapper{
				MetricData: item.MetricData,
				identifier: identifier,
				order:      order,
			})

			for _, mv := range item.MetricValues {
				// Record ordinals in the final result
				v, ok := mv.Metadata[identifier]
				if ok && v != "" {
					resourceOrdinal[v] = ordinal
					ordinal++
				}
			}
		}

		// Add every unique identifier value to the set
		for _, mv := range item.MetricValues {
			v, ok := mv.Metadata[identifier]
			if ok && v != "" {
				resourceSet[v] = true
			}
		}
	}

	var resourceList []string
	for k := range resourceSet {
		resourceList = append(resourceList, k)
	}
	sort.Strings(resourceList)

	// Fill resource-ordinal map with resources never present in the target, and give them ordinals.
	for _, r := range resourceList {
		if _, ok := resourceOrdinal[r]; !ok {
			resourceOrdinal[r] = ordinal
			ordinal++
		}
	}

	// Sort metrics
	for i, item := range raw.Results {
		if item.MetricType != monitoring.MetricTypeVector || item.Error != "" {
			continue
		}

		sorted := make([]monitoring.MetricValue, len(resourceList))
		for _, mv := range item.MetricValues {
			v, ok := mv.Metadata[identifier]
			if ok && v != "" {
				ordinal, _ := resourceOrdinal[v]
				sorted[ordinal] = mv
			}
		}
		raw.Results[i].MetricValues = sorted
	}

	raw.CurrentPage = 1
	raw.TotalPages = 1
	raw.TotalItems = len(resourceList)
	return raw
}

func (raw *Metrics) Page(page, limit int) *Metrics {
	if page < 1 || limit < 1 || len(raw.Results) == 0 {
		return raw
	}

	start := (page - 1) * limit
	end := page * limit

	for i, item := range raw.Results {
		if item.MetricType != monitoring.MetricTypeVector || item.Error != "" {
			continue
		}

		total := len(item.MetricValues)
		if start >= total {
			raw.Results[i].MetricValues = nil
			continue
		}
		if end >= total {
			end = total
		}

		raw.Results[i].MetricValues = item.MetricValues[start:end]
	}

	raw.CurrentPage = page
	raw.TotalPages = int(math.Ceil(float64(raw.TotalItems) / float64(limit)))
	return raw
}

type SortOptions struct {
	target     string
	identifier string
	order      string
	page       int
	limit      int
}

func (opt *SortOptions) ShouldSort() bool {
	return opt.target != "" && opt.identifier != ""
}

func parseSortOptions(req *restful.Request) *SortOptions {
	opt := &SortOptions{
		page:  DefaultPage,
		limit: DefaultLimit,
		order: DefaultOrder,
	}

	target := req.QueryParameter("sort_metric")
	order := req.QueryParameter("sort_type")
	page := req.QueryParameter("page")
	limit := req.QueryParameter("limit")

	if target == "" {
		return opt
	}
	opt.target = target

	if page != "" {
		pageVal, err := strconv.Atoi(page)
		if err == nil {
			opt.page = pageVal
		}
	}

	switch order {
	case OrderAscending:
	case OrderDescending:
		opt.order = order
	}

	if limit != "" {
		limitVal, err := strconv.Atoi(limit)
		if err == nil {
			opt.limit = limitVal
		}
	}

	return opt
}
