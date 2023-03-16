package main

import (
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

//func init() {
//	tracer.Start(
//		tracer.WithService("test-svc"), tracer.WithEnv("test-env"))
//	defer tracer.Stop()
//}

func main() {
	tracer.Start()
	defer tracer.Stop()

	r := gin.New()

	r.Use(gintrace.Middleware("test-svc3-new"))

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	r.GET("/getUser/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, map[string]interface{}{
			"name":   name,
			"status": http.StatusOK,
		})
	})

	r.GET("/getLocal/:address", func(c *gin.Context) {
		address := c.Param("address")
		if address == "" {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status": http.StatusInternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"address": address,
			"status":  http.StatusOK,
		})
	})

	r.Run(":8080")

}
