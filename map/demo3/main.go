package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	name string `json:"name"`
	age int `json:"age"`
	address string `json:"address"`
}

func main() {
	m := map[string]string{}
	m["name"] = "jerry"

	val,ok := m["name1"]
	if val == "jerry"{
		fmt.Println("val:",val)
	}else {
		fmt.Println("false val:",val)
	}

	fmt.Println("ok:",ok)

	users := []user{}

	u := users

	users = append(users, u...)

	data,err := json.Marshal(&users)
	if err !=nil{
		fmt.Println("err:",err)
	}

	fmt.Println("data:",string(data))
}
