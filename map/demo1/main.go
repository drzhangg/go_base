package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Username    string        `json:"username"`
	RoleBinding []Rolebinding `json:"roleBindings"`
}

type Rolebinding struct {
	Zone      string `json:"zone"`
	Namespace string `json:"namespace"`
	Role      string `json:"role"`
}

func main() {
	data, err := ioutil.ReadFile("/Users/zhang/drzhang/go/go_base_util/file/demo5/rolebinding.json")
	if err != nil {
		fmt.Println(err)
	}

	ds := Data{}
	json.Unmarshal(data,&ds)

	fmt.Println("ds::", len(ds.RoleBinding))


	var zoneMap = make(map[string]string)
	for _, v := range ds.RoleBinding{
		_,ok := zoneMap[v.Zone + ";" + v.Namespace]
		if !ok{
			zoneMap[v.Zone + ";" + v.Namespace] = v.Role
		}
		//zoneMap[v.Zone] = struct{}{}
	}
	fmt.Println(zoneMap)

	//fmt.Println("data::", string(data))
}
