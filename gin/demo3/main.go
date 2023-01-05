package main

import "github.com/gin-gonic/gin"

type User struct {
	BaseQuantity `json:",omitempty"`
	Name string
}

type BaseQuantity map[string]map[string]int

func main() {
	r := gin.Default()

	r.GET("/test", test)

	r.Run(":8282")
}

func test(c *gin.Context) {
	m := make(map[string]map[string]int)

	m1 := make(map[string]int)
	m1["dev"] = 1
	m1["poc"] = 3
	m1["sos"] = 4
	m["vm"] = m1

	m2 := make(map[string]int)
	m2["dev"] = 18
	m2["poc"] = 33
	m["vd"] = m2
	u := User{
		BaseQuantity: m,
		Name:         "zhang",
	}

	c.JSON(200,u)
}
