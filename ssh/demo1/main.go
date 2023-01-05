package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {

	client,err := ssh.Dial("tcp","150.158.87.137:22",&ssh.ClientConfig{
		Config:            ssh.Config{},
		User:              "root",
		Auth:              []ssh.AuthMethod{ssh.Password("zjh94264,")},
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("ssh dial error: %s",err)
	}

	session,err := client.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("iptables -L -n -t nat"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}
