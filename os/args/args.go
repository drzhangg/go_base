package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	length := len(args)

	fmt.Println(length)
	fmt.Println(args[0])


	base := filepath.Base(args[0])
	fmt.Println(base)

	name,_ := os.Hostname()
	fmt.Println(name)
}
