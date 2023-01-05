package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "qweradazcs"

	base64S := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(base64S)
}
