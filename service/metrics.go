package service

import "github.com/prometheus/client_golang/prometheus"

var (
	Count = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "count",
			Help: "total count",
		},
		[]string{"action"},
	)
	Latency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "latency",
			Help: "latency (unit: second)",
		},
		[]string{"action"},
	)
)

func init() {
	prometheus.MustRegister(Count)
	prometheus.MustRegister(Latency)
}
