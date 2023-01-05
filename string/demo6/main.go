package main

import (
	"fmt"
	"github.com/moby/term"
	"os"
)

func main() {
	var buffer [512]byte

	//os.Stdin.Read()
	fmt.Println(os.Stdout.Fd())

	n, err := os.Stdin.Read(buffer[:])
	if err != nil {

		fmt.Println("read error:", err)
		return

	}
	fmt.Println(os.Stdin.Fd())

	state,err := term.MakeRaw(os.Stdin.Fd())
	if err !=nil{
		fmt.Println("err:",err)
	}

	err = term.RestoreTerminal(os.Stdin.Fd(),state)
	if err !=nil{
		fmt.Println("err:",err)
	}


	fmt.Println("count:", n, ", msg:", string(buffer[:]))

}
