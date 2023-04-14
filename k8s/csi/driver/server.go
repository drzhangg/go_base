package driver

import (
	"context"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"google.golang.org/grpc"
	"k8s.io/klog"
	"net"
	"os"
	"strings"
)

var (
	driverName = "zjh.csi.com"
	version    = "1.0.0"
)

func NewDriver(nodeId, endpoint string) *Driver {
	return &Driver{
		nodeID:   nodeId,
		endpoint: endpoint,
	}
}

type Driver struct {
	nodeID   string
	endpoint string
}

func (d *Driver) Run() {
	ctl := NewControllerServer()
	identity := NewIdentityServer()
	node := NewNodeServer(d.nodeID)

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logGRPC),
	}

	srv := grpc.NewServer(opts...)

	csi.RegisterControllerServer(srv,ctl)
	csi.RegisterIdentityServer(srv,identity)
	csi.RegisterNodeServer(srv,node)


	proto,addr,err := ParseEndpoint(d.endpoint)
	klog.V(4).Infof("protocol: %s,addr: %s", proto, addr)
	if err != nil {
		klog.Fatal(err.Error())
	}

	if proto == "unix" {
		addr = "/" + addr
		if err := os.Remove(addr);err != nil && !os.IsNotExist(err){
			klog.Fatalf("Failed to remove %s, error: %s", addr, err.Error())
		}
	}

	listener,err := net.Listen(proto,addr)
	if err!= nil {
		klog.Fatalf("Failed to listen: %v", err)
	}

	srv.Serve(listener)

}

func logGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	klog.V(4).Infof("GRPC call: %s", info.FullMethod)
	klog.V(4).Infof("GRPC request: %s", protosanitizer.StripSecrets(req))

	resp, err := handler(ctx, req)
	if err != nil {
		klog.Errorf("GRPC error: %v", err)
	}
	klog.V(4).Infof("GRPC response: %s", protosanitizer.StripSecrets(resp))

	return resp, err
}

func ParseEndpoint(endpoint string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(endpoint), "unix://") || strings.HasPrefix(strings.ToLower(endpoint), "tcp://") {
		s := strings.SplitN(endpoint, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("Invalid endpoint: %v", endpoint)
}
