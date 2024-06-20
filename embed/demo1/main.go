package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var f embed.FS

func main() {

	data,_ := f.ReadFile("hello.txt")
	//fmt.Println(data)
	fmt.Println(string(data))

}
