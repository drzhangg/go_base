package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	r, _,_ := reader.ReadLine()

	os.Stdout.WriteString("stdout:"+string(r))
	os.Stderr.WriteString("stderr:"+string(r))

}
