package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//fmt.Print("\033[H\033[2J")

	//path := "/usr/local/bin/kubectl"
	path := "/usr/local/bin"

	s := filepath.Base(path)
	fmt.Println(s)
}
