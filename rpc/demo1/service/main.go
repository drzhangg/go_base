package main

import (
	"log"
	"net"
	"net/rpc"
)

type User struct {
	Name string
	Age int
}

type HelloService struct {
}

func (p *HelloService) Hello(request string, user *User) error {
	user = &User{
		Name: "jerry",
		Age:  26,
	}
	return nil
}

func main() {
	rpc.RegisterName("HelloService",new(HelloService))

	listener,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("ListenTCP error:",err)
	}

	conn,err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:",err)
	}

	rpc.ServeConn(conn)
}
