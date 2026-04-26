package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SERVICE = "blog"

var (
	requestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_counter",
			Help: "Number of requests received",
		},
		[]string{"service", "interface"},
	)

	// Histogram 记录延迟分布，支持分位数聚合
	requestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_ms",
			Help:    "Request duration in milliseconds",
			Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000, 2500, 5000},
		},
		[]string{"service", "interface"},
	)
)

func Metric() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		begin := time.Now()
		ctx.Next()
		ifc := mappingUrl(ctx)
		requestCounter.WithLabelValues(SERVICE, ifc).Inc()
		requestDuration.WithLabelValues(SERVICE, ifc).Observe(float64(time.Since(begin).Milliseconds()))
	}
}

var restfulMapping = map[string]string{"uid": ":uid", "bid": ":bid"}

func mappingUrl(ctx *gin.Context) string {
	url := ctx.Request.URL.Path
	for _, p := range ctx.Params {
		if value, exists := restfulMapping[p.Key]; exists {
			url = strings.Replace(url, p.Value, value, 1)
		}
	}
	return url
}
