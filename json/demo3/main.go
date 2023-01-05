package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Phone string   `json:"phone"`
	Info  UserInfo `json:"info"`
}

type UserInfo map[string]int

func main() {
	m := make(map[string]int)
	u := User{}
	m["name"] = 1
	m["age"] = 26
	m["phone"] = 110

	marshalData(m,&u.Info)
	fmt.Println(u)

}

func marshalData(m map[string]int, b *UserInfo) {
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &b)
	if err != nil {
		fmt.Println(err)
	}
}
