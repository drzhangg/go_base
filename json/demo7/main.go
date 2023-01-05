package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var m = make(map[string]Data)

type Demo struct {
	Data   `json:"data"`
	Status `json:"status"`
}

type Data struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Status struct {
	Message      string `json:"message"`
	DeployStatus string `json:"deployStatus"`
	InternalName string `json:"internalName"`
	CommitId     string `json:"commitId"`
}

func main() {
	file, err := ioutil.ReadFile("/Users/zhang/drzhang/demo/go/go_base/json/demo7/demo.json")
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("file:", string(file))

	var i interface{}
	err = json.Unmarshal(file,&i)
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println("i:::",i)

	sts := Demo{}
	err = json.Unmarshal(file,&sts)
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println("sts:::",sts)


	//m := make(map[string]interface{})
	//json.Unmarshal(file, &m)
	//
	//
	//sts,err := json.Marshal(m["status"])
	//
	//status := Status{}
	//err = json.Unmarshal(sts,&status)
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("s:",status)
	//
	//status.Message = "hello world"
	//status.DeployStatus = "failed"
	//
	//m["status"] = status
	//
	//fmt.Println(m)



	//d := Demo{
	//	Data{
	//		Name: "jerry",
	//	},
	//}

	//err = json.Unmarshal(file, &d)
	//if err != nil {
	//	fmt.Println("unmarshal err:", err)
	//}
	//fmt.Println(d)

}
