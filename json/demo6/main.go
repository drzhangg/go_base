package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//s := "{\"status\":\"error\",\"message\":\"nothing created.\"}\n{\"status\":\"error\",\"message\":\"Failed to mount /sdata04.\"}"
	s := "{\"status\":\"error\",\"message\":\"nothing created.\"}"


	var s1 string
	err := json.Unmarshal([]byte(s),&s1)
	if err != nil {
		fmt.Println("errr1:",err)
	}
	fmt.Println("s1:",s1)

	data,err := json.Marshal(s)
	if err != nil {
		fmt.Println("err1:",err)
	}

	fmt.Println(string(data))

	var str interface{}
	err = json.Unmarshal(data,&str)
	if err != nil {
		fmt.Println("err1:",err)
	}
	fmt.Println(str)

	type Data struct {
		Status string `json:"status"`
		Message string `json:"message"`
	}

	dd,err:=json.Marshal(str)
	if err != nil {
		fmt.Println("err3:",err)
	}

	data1 := Data{}
	err = json.Unmarshal(dd,&data1)
	if err != nil {
		fmt.Println("err4:",err)
	}

	//fmt.Println(str)
	fmt.Println(data1)
}
