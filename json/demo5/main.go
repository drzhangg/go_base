package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleType string `json:"roleType"`
	Age      int    `json:"age"`
}

func main() {

	data, err := ioutil.ReadFile("/Users/zhang/drzhang/demo/go/go_base/json/demo5/data.json")
	//data,err := ioutil.ReadFile("./go_base/json/demo5/data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("data::", string(data))

	d := Data{}
	d.Age = 26
	err = json.Unmarshal(data,&d)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("d:",d)



	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	m1 := make(map[string]string)

	d1, err := json.Marshal(m["credentials"])
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(d1, &m1)
	fmt.Println("m1:", m1)

}
