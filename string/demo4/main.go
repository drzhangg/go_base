package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/moby/term"
	"os"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	//strBytes, _, err := inputReader.ReadLine()

	//os.Stdin.WriteString()
	//term.GetFdInfo(inputReader)
	inFd, _ := term.GetFdInfo(inputReader)

	state, err := term.MakeRaw(inFd)

	term.RestoreTerminal(inFd, state)


	//"{\\\"Op\\\":\\\"stdout\\\",\\\"Data\\\":\\\"\\\\u001b[1;32mhello\\\\u001b[m  \\\\u001b[1;34mlogs\\\\u001b[m\\\\r\\\\n\\\",\\\"SessionID\\\":\\\"\\\",\\\"Rows\\\":0,\\\"Cols\\\":0}"
	//msg := "{\"Op\":\"stdout\",\"Data\":\"\\u001b[1;32mhello\\u001b[m  \\u001b[1;34mlogs\\u001b[m\\r\\n/app # \\u001b[6n\",\"SessionID\":\"\",\"Rows\":0,\"Cols\":0}"
	msg := "{\\\"Op\\\":\\\"stdout\\\",\\\"Data\\\":\\\"\\\\u001b[1;32mhello\\\\u001b[m  \\\\u001b[1;34mlogs\\\\u001b[m\\\\r\\\\n\\\",\\\"SessionID\\\":\\\"\\\",\\\"Rows\\\":0,\\\"Cols\\\":0}"
	fmt.Println([]byte(msg))
	m := make(map[string]string)
	err = json.Unmarshal([]byte(msg), &m)
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println("mmmmmm:", m)
	//fmt.Println("mmmmmm:", d)

	//fmt.Println("{\"Op\":\"stdout\",\"Data\":\"\u001b[1;32mhello\u001b[m  \u001b[1;34mlogs\u001b[m\r\n/app # \u001b[6n\",\"SessionID\":\"\",\"Rows\":0,\"Cols\":0}")
	//
	//s := "a[\"{\"Op\":\"stdout\",\"Data\":\"/app # \u001b[6n\",\"SessionID\":\"\",\"Rows\":0,\"Cols\":0}\"]"
	////s1 := "/app # \u001b[6n\"
	//fmt.Printf("/app # \u001b[6n")

}
