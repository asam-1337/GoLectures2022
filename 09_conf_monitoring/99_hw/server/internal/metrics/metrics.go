package metrics

import (
	"fmt"
	"github.com/labstack/echo/v4"
	prom "github.com/prometheus/client_golang/prometheus"
)

const (
	KeyMetrics = "custom_metric"
)

var (
	RequestId = prom.NewCounter(prom.CounterOpts{
		Name: "total_requests",
	})
	TotalReqCnt = prom.NewCounterVec(prom.CounterOpts{
		Name: "request_status_url",
	}, []string{"status", "path"})
	TimingDur = prom.NewHistogram(prom.HistogramOpts{})
)

type Metrics struct {
	RequestId   int
	TotalReqCnt prom.Counter
	Hits        *prom.CounterVec
	TimingDur   prom.Histogram
}

func NewMetrics() *Metrics {
	m := &Metrics{
		TotalReqCnt: prom.NewCounter(prom.CounterOpts{
			Name: "total_requests",
			Help: "Counts executions of my handler function.",
		}),
		Hits: prom.NewCounterVec(prom.CounterOpts{
			Name: "request_status_url",
		}, []string{"status", "path"}),
		TimingDur: prom.NewHistogram(prom.HistogramOpts{
			Name:    "timings",
			Buckets: prom.DefBuckets,
		}),
	}
	err := prom.Register(m.TotalReqCnt)
	if err != nil {
		fmt.Printf("cant register metric: %s", err.Error())
		return nil
	}

	err = prom.Register(m.Hits)
	if err != nil {
		fmt.Printf("cant register metric: %s", err.Error())
		return nil
	}

	err = prom.Register(m.TimingDur)
	if err != nil {
		fmt.Printf("cant register metric: %s", err.Error())
		return nil
	}
	return m
}

func (m *Metrics) MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("OK")
		m.TotalReqCnt.Inc()
		m.Hits.WithLabelValues("200", c.Path()).Inc()
		c.Set(KeyMetrics, m)
		m.RequestId++
		return next(c)
	}
}
