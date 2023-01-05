package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func RemoteShell(ws *websocket.Conn) {
	bScanner := bufio.NewScanner(ws)
	currentWorkingDir, _ := os.Getwd()
	fmt.Println("current working dir", currentWorkingDir)

	for bScanner.Scan() {
		// parse command
		cmd := bScanner.Text()
		fmt.Println(cmd)
		cmdItems := strings.Split(cmd, " ")
		cmdName := cmdItems[0]

		var cmdArgs []string
		if len(cmdItems) >= 2 {
			cmdArgs = cmdItems[1:]
		}

		// execute command
		cmdOutput := bytes.NewBuffer(nil)
		cmdExec := exec.Command(cmdName, cmdArgs...)
		cmdExec.Dir = currentWorkingDir
		cmdExec.Stdout = cmdOutput
		cmdExec.Stderr = cmdOutput
		err := cmdExec.Run()
		if err != nil {
			fmt.Println(err)
			ws.Write([]byte(err.Error()))
		} else {
			ws.Write(cmdOutput.Bytes())
		}
	}
}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "0.0.0.0", "host to listen")
	flag.IntVar(&port, "port", 9001, "port to listen")
	flag.Parse()

	//handler
	http.Handle("/remote/shell", websocket.Handler(RemoteShell))

	//listen
	endPoint := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(endPoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}