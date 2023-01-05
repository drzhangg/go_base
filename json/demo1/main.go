package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type raw struct {
	Raw []byte `json:"-"`
}

type Data struct {
	Demo
	Name    string `json:"name"`
	Age     int    `json:"age"`
	BeginAt string `json:"begin_at"`
}

type Demo struct {
	Example map[string]interface{} `json:"-"`
}

func main() {

	var r raw

	bytes, _ := ioutil.ReadFile("/Users/zhang/drzhang/demo/go/go_base/json/demo1/user.json")
	//fmt.Println("bbb:", string(bytes))

	r.Raw = bytes

	var data Data
	err := json.Unmarshal(r.Raw, &data)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(data.BeginAt)

}
