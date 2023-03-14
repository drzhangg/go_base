package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	}, []string{"method", "path", "status"})
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

type User struct {
	Gender  int    `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func hello(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		name = "jerry"
	}

	c.JSON(http.StatusOK, map[string]string{"name": name})
}


func getEnv(c *gin.Context) {

	vmCrd := os.Getenv("VM_CRD")
	c.JSON(http.StatusOK, map[string]interface{}{
		"vm_crd": vmCrd,
	})
}

func command(c *gin.Context) {

	finfos, err := ioutil.ReadDir("/etc/config.conf")
	if err != nil {
		fmt.Println("err:", err)
	}

	for _, v := range finfos {
		if !v.IsDir() {
			data, err := ioutil.ReadFile(v.Name())
			if err != nil {
				fmt.Println("err:", err)
			}
			fmt.Println("data::", string(data))
		} else {
			fmt.Println("没有文件输出")
		}
	}

}

func createUser(c *gin.Context) {
	name := c.Param("username")

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln("err:", err)
	}

	var genderMap = make(map[int]string)
	genderMap[0] = "男"
	genderMap[1] = "女"

	c.JSON(http.StatusOK, map[string]interface{}{
		"name":    name,
		"age":     strconv.Itoa(user.Age) + "岁",
		"gender":  genderMap[user.Gender],
		"address": user.Address,
	})
}

func deleteUser(c *gin.Context) {
	username := c.Param("username")
	if username == "jerry"{
		c.JSON(200,"delete user success")
		return
	}
	c.JSON(500,"delete user failed")
}

func main() {
	r := gin.Default()

	// 注册指标路由
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 注册中间件
	r.Use(monitorMetrics())

	r.GET("/hello/:name", hello)

	r.POST("/user/:username", createUser)

	r.GET("/getEnv", getEnv)

	r.GET("/command", command)

	r.DELETE("/delete/:username",deleteUser)

	r.Run(":8181")
}
