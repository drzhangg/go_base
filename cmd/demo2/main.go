package main

import (
	"fmt"
	"os"
)

func main() {

	//inFd, isTerminal := term.GetFdInfo(t.In)

	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, terminationSignals...)
	//defer func() {
	//	signal.Stop(ch)
	//	close(ch)
	//}()


	f,err := os.Open("/dev/tty")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f.Fd())
	fmt.Println(f.Stat())
}
