package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var res = schema.GroupVersionResource{
	Group:    "drzhangg.my.domain",
	Version:  "v1beta1",
	Resource: "Frigate",
}

type Server struct {
	client Interface
}

type Interface interface {
	kubernetes.Interface
	dynamic.Interface
}

func (s *Server) getDns() {
	s.client.Resource(res).List(context.TODO(), v1.ListOptions{})
}

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config.conf")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	clientset = clientset
	//clientset.
}
