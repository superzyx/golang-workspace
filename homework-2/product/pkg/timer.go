package web

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const metricName = "httpserver"

func Register() {
	err := prometheus.Register(histogram)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("ready")
	}
}

var(
	histogram = NewHistogramVec(metricName, "time spend")
)

type ExecutionMethod interface {
	NewHistogramVec(namespace, help string) *prometheus.HistogramVec
}

func NewHistogramVec(namespace, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name: "execution_latency_time",
			Buckets: prometheus.ExponentialBuckets(0.001,2,16),
			Help: help},
		[]string{"step"},)
}

func NewMetrics() FilterBuilder{
	return func(f Filter) Filter {
		return func(c *Context) {
			et := NewExecutionTimer()
			f(c)
			et.data.WithLabelValues("total").Observe(time.Now().Sub(et.start).Seconds())
		}
	}
}

func NewExecutionTimer() *ExecutionTimer{
	t := time.Now()
	return &ExecutionTimer{
		histogram,
		t,
		t,
	}
}


//ExecutionTimer measures execution time of a computation,split into major step
type ExecutionTimer struct {
	data *prometheus.HistogramVec
	start time.Time
	end time.Time
}