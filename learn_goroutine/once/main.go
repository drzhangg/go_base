package main

import (
	"net"
	"sync"
	"time"
)

var connMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	// 返回已经创建好的连接
	if conn != nil {
		return conn
	}

	// 创建连接
	conn,_ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

func main() {
	conn := getConn()
	if conn == nil {
		panic("conn is nil")
	}
}
