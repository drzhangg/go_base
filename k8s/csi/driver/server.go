package driver

import (
	"google.golang.org/grpc"
	"sync"
)

func NewNonBlockingGRPCServer() *nonBlockingGRPCServer {
	return &nonBlockingGRPCServer{}
}

type nonBlockingGRPCServer struct {
	wg      sync.WaitGroup
	server  *grpc.Server
	cleanup func()
}

func (*nonBlockingGRPCServer)serve()  {
	
}
