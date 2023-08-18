package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	user := r.Group("user", middleware1(), middleware2())
	user.GET("/get", func(c *gin.Context) {
		c.JSON(200, map[string]string{"name": "jerry"})
	})

	r.Run(":9191")
}

func middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware1 start")
		c.Next()
		fmt.Println("middleware1 end")
	}
}

func middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware2 start")
		//c.Next()
		fmt.Println("middleware2 end")
	}
}
