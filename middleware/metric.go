package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strings"
	"time"
)

const SERVICE = "blog"

var (
	// Counter和Gauge指标定义
	requestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_counter",
			Help: "Number of requests received",
		},
		[]string{"service", "interface"},
	)

	requestTimer = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "request_timer",
			Help: "Time spent processing request",
		},
		[]string{"service", "interface"},
	)
)

// Metric 中间件函数
func Metric() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		begin := time.Now()
		ctx.Next()
		ifc := mappingUrl(ctx)
		requestCounter.WithLabelValues(SERVICE, ifc).Inc()
		requestTimer.WithLabelValues(SERVICE, ifc).Set(float64(time.Since(begin).Milliseconds()))
	}
}

var restfulMapping = map[string]string{"uid": ":uid", "bid": ":bid"}

// 映射URL到标准restful格式
func mappingUrl(ctx *gin.Context) string {
	url := ctx.Request.URL.Path
	for _, p := range ctx.Params {
		if value, exists := restfulMapping[p.Key]; exists {
			url = strings.Replace(url, p.Value, value, 1)
		}
	}
	return url
}
