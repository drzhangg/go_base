package main

import (
	"encoding/json"
	"fmt"
)

type b struct {
	Name   string `json:"name"`
	School string `json:"school"`
	Grade  int    `json:"grade"`
}

func main() {
	a := map[string]interface{}{"name": "sam", "age": 18}

	var b1 b

	data, _ := json.Marshal(a)

	json.Unmarshal(data, &b1)
	fmt.Println(b1)
}
