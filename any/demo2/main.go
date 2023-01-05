package main

import "fmt"

func main() {
	m := map[string]interface{}{"10.10.10.10":"{\"status\":\"error\",\"message\":\"nothing created.\"}"}

	ip := "10.10.10.10"
	result,ok := m[ip].(string)
	if !ok {
		fmt.Println("not ok")
	}

	fmt.Println("result:",result)
}
