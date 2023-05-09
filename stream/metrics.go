package stream

import (
	"github.com/blocknative/dreamboat/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type StreamMetrics struct {
	StreamRecvCounter *prometheus.CounterVec
	Timing            *prometheus.HistogramVec
}

func (s *Client) initMetrics() {
	s.m.StreamRecvCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dreamboat",
		Subsystem: "stream",
		Name:      "recvcount",
		Help:      "Number of blocks received from stream.",
	}, []string{"type"})

	s.m.Timing = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "dreamboat",
		Subsystem: "stream",
		Name:      "timing",
		Help:      "Duration of requests per function",
	}, []string{"function", "type"})
}

func (s *Client) AttachMetrics(m *metrics.Metrics) {
	m.Register(s.m.StreamRecvCounter)
	m.Register(s.m.Timing)
}