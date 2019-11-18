package collector

import "github.com/prometheus/client_golang/prometheus"

// StatsCollector defines an interface for collecting specific stats
// from a nsqd exported stats data.
type StatsCollector interface {
	set(s *stats)
	Collect(out chan<- prometheus.Metric)
	Describe(ch chan<- *prometheus.Desc)
	reset()
}
