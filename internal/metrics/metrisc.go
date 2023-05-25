package metrics

import (
	"github.com/rcrowley/go-metrics"
)

const (
	GetWithMaxTime = "get_max"
	GetWithMinTime = "get_min"
	GetByURL       = "get_by_url"
)

func Init() {
	metrics.Unregister(GetWithMaxTime)
	metrics.MustRegister(GetWithMaxTime, metrics.NewCounter())

	metrics.Unregister(GetWithMinTime)
	metrics.MustRegister(GetWithMinTime, metrics.NewCounter())

	metrics.Unregister(GetByURL)
	metrics.MustRegister(GetByURL, metrics.NewCounter())
}

// IncCounter increments counter by name
func IncCounter(name string) {
	counter := metrics.Get(name).(metrics.Counter)
	counter.Inc(1)
}

// GetCount returns a value from counter by name
func GetCount(name string) int64 {
	c := metrics.Get(name).(metrics.Counter)
	return c.Count()
}

// ClearCount clears counter by name
func ClearCount(name string) {
	metrics.Get(name).(metrics.Counter).Clear()
}

// GetCounterAndClear returns a value and clears counter
func GetCounterAndClear(name string) int64 {
	c := metrics.Get(name).(metrics.Counter)
	count := c.Count()
	c.Clear()
	return count
}
