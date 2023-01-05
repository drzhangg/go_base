package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn,err := net.Dial("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout,conn)
		log.Println("done")

		done <- struct{}{} // 向主goroutine发出信号
	}()
	mustCopy(conn,os.Stdin)
	conn.Close()
	<- done //等待后台goroutine完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _,err := io.Copy(dst,src);err != nil {
		log.Fatal(err)
	}
}
