package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	ca = cache.New(time.Minute*10, time.Minute*15)
)

func main() {

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {

		key := "name"

		name, ok := ca.Get(key)
		if ok {
			fmt.Println("get ok")
			c.JSON(200, "get name:" + name.(string))
			return
		}
		fmt.Println("get failed")

		ca.Set(key,"bob",time.Second * 20)
		//ca.SetDefault(key, "jerry")
		c.JSON(200, "get name failed")
	})

	r.Run(":8989")

}

//func init() {
//	ca.Set("name","tom",time.Second * 5)
//}
