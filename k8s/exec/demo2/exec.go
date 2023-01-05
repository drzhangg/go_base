package main

import (
	"os"
	"os/exec"
)

func main() {

	//a1 := "--context"
	//a2 := "tencent"
	a3 := "exec"
	a4 := "-it"
	a5 := "-n"
	a6 := "default"
	a7 := "nginx-test-6c88c44876-vf9vs"
	a8 := "-c"
	a9 := "nginx"
	a10 := "--"
	a11 := "sh"
	a12 := "-c"
	a13 := "command -v bash >/dev/null && exec bash || exec sh"
	cmd := exec.Command("/usr/local/bin/kubectl", a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)

	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	cmd.Run()
}
