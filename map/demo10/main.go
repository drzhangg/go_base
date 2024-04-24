package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {

	u := User{
		Name:    "jerry",
		Age:     30,
		Address: "china wuhan",
	}

	data,_ :=json.Marshal(u)


	m := make(map[string]interface{})

	err := json.Unmarshal(data,&m)
	if err != nil {
		fmt.Println("err:",err)
	}

	fmt.Println(m)

	m1 := make(map[string]string)
	for key, value := range m {
		switch v := value.(type) {
		case string:
			m1[key] = v
		default:
			m1[key] = fmt.Sprintf("%v", v)
		}
	}
	fmt.Println("m1:",m1)


	//m1 := make(map[string][]string)
	//m1["one"] = []string{}
	//m1["two"] = []string{}
	//m1["three"] = []string{}
	//
	//a1 := "one"
	////a2 := "two"
	////a3 := "three"
	//
	////for{
	//m1[a1] = append(m1[a1], "1")
	//m1[a1] = append(m1[a1], "2")
	//m1[a1] = append(m1[a1], "3")
	////}
	//
	//fmt.Println(m1)

}
