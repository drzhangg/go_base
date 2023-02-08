package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

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

	r.GET("/hello/:name", hello)

	r.POST("/user/:username", createUser)

	r.GET("/getEnv", getEnv)

	r.GET("/command", command)

	r.DELETE("/delete/:username",deleteUser)

	r.Run(":8181")
}
