package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	}, []string{"method", "path", "status"})

	//httpRequestsTime = promauto.NewCounterVec(prometheus.CounterOpts{
	//	Name: "http_requests_total",
	//	Help: "Total number of HTTP requests.",
	//},[]string{"method", "path", "status"})
)

func monitorMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		//start := time.Now()

		// 调用处理程序
		c.Next()

		// 获得响应状态码
		status := strconv.Itoa(c.Writer.Status())


		// 记录指标
		httpRequestsTotal.WithLabelValues(c.Request.Method, c.Request.URL.Path, status).Inc()
		//httpRequestsTime.WithLabelValues(c.Request.Method,c.Request.URL.Path,start.Sub(time.Now()).String()).Add()
	}
}

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 注册指标路由
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 注册中间件
	r.Use(monitorMetrics())

	//r.Group("/metrics")

	// 启动Gin引擎
	r.Run(":8282")
}
