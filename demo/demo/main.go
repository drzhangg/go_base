package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	f := fmt.Sprintf(`{"Authorization": "Internal %s"}`,"jerry")
	fmt.Println(f)


	d,_ := base64.StdEncoding.DecodeString("emhhbmcK")
	fmt.Println(string(d))
}
