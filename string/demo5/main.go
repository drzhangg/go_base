package main

import (
	"fmt"
	"github.com/creack/pty"
	"github.com/moby/term"
	"log"
)

func main() {
	_, tty, err := pty.Open()
	if err != nil {
		//log.Fatalf()
		log.Fatalf("error creating pty: %v", err)
	}

	tty.Read([]byte("sadas"))

	inFd, isTerminal := term.GetFdInfo(tty)
	fmt.Println("isTerminal:", isTerminal)

	state, _ := term.MakeRaw(inFd)

	err = term.RestoreTerminal(inFd, state)
	if err != nil {
		fmt.Println("err:", err)
	}

	//fd := os.Stdin.Fd()
	//if term.IsTerminal(fd) {
	//	ws, err := term.GetWinsize(fd)
	//	if err != nil {
	//		log.Fatalf("term.GetWinsize: %s", err)
	//	}
	//	log.Printf("%d:%d\n", ws.Height, ws.Width)
	//}
}
