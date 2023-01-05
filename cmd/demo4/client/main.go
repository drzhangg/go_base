package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"os"
)

func main() {
	var origin string
	var url string
	flag.StringVar(&origin, "origin", "", "websocket origin")
	flag.StringVar(&url, "url", "", "websocket remote url")
	flag.Parse()

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
		return
	}

	buffer := make([]byte, 40960)
	bScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("&gt; ")
	for bScanner.Scan() {
		line := bScanner.Text()
		ws.Write([]byte(line + "\r\n"))
		num, err := ws.Read(buffer)
		if err != nil {
			ws.Close()
			return
		}
		fmt.Print(string(buffer[:num]))
		//fmt.Print("&gt; ")
	}
}
