package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type data struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func main() {
	//s := `{\"status\":\"success\",\"message\":\"domain sdev\suimedia54 added.\"}`
	s1 := "{\"status\":\"success\",\"message\":\"domain sdevsuimedia54 added.\"}"
	//s := `{"status":"success","message":"domain sdev/suimedia54 added."}`

	s := strings.ReplaceAll(s1,"\\","/")

	//m := make(map[string]interface{})

	m := data{}

	err := json.Unmarshal([]byte(s),&m)
	if err != nil {
		fmt.Println("unmarshal err:",err)
	}

	nm := strings.ReplaceAll(m.Message,"/","\\")
	m.Message = nm
	fmt.Println(m)
}
