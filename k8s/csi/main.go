package main

import (
	"csi/driver"
	"flag"
	"k8s.io/klog"
)

var (
	endpoint string
	nodeId   string
)

func main() {
	flag.StringVar(&endpoint, "endpoint", "", "CSI Endpoint")
	flag.StringVar(&nodeId, "nodeid", "", "node id")

	klog.InitFlags(nil)
	flag.Parse()

	d := driver.NewDriver(nodeId, endpoint)
	d.Run()
}
