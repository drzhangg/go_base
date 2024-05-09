package main

import (
	"fmt"
	"strings"
)

func main() {
  operation := "/api.platform.role.v1.Role/List"

  arr := strings.Split(operation,"/")

  fmt.Println(arr[1])

  str1 := "list"
  str2 := "ListEvent"
  fmt.Println(strings.Contains(strings.ToLower(str2),str1))

}
