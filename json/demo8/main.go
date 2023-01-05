package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := "{\"status\":\"error\",\"message\":\"nothing created.\"}"
	//s := "{\"status\":\"error\",\"message\":\"nothing created.\"}\n{\"status\":\"error\",\"message\":\"Failed to mount /sdata04.\"}"
	type Data struct {
		Status string `json:"status"`
		Message string `json:"message"`
	}


	var str Data
	err := json.Unmarshal([]byte(s),&str)
	if err != nil {
		fmt.Println("err1:",err)
	}
	fmt.Println(str)

}
