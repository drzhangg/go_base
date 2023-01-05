package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
}

func main() {
	m := make(map[string]int)

	d, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}

	var u user

	err = json.Unmarshal(d, &u)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("user:", u)
}
