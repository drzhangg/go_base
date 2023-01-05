package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func main() {
	file,err := ioutil.ReadFile("/Users/zhang/drzhang/demo/go/go_base/k8s/unstructured/demo1/test.yaml")
	if err !=nil{
		fmt.Println("read file err:",err)
	}

	fmt.Println(string(file))

	obj := map[string]interface{}{}

	err = json.Unmarshal(file,&obj)
	if err !=nil{
		fmt.Println("unmarshal err:",err)
	}

	type status struct {
		State string `json:"state"`
	}

	obj["status"] = status{State: "failed"}


	un := &unstructured.Unstructured{
		Object: obj,
	}

	fmt.Println(un)
}
