package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

var (
	url = "/v1/service_instance/%s/service_binding/%s"
)

func main() {
	host := "127.0.0.1"
	port := "8080"

	instanceId := uuid.New().String()
	bindingId := uuid.New().String()
	fmt.Println(fmt.Sprintf(host+":"+port+url,instanceId,bindingId))

	m := make(map[string]interface{})
	m["app_id"] = "asdasd12312"
	m["type"] = "pod"
	m["param"] = struct {
		Name string `json:"name"`
		Address string `json:"address"`
	}{
		Name: "jerry",
		Address: "shanghai pudong",
	}
	m1 := make(map[string]interface{})
	m1["parameters"] = m

	data,err := json.Marshal(&m1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
