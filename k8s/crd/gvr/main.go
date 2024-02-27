package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	h := os.Getenv("HOME")

	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("BuildConfigFromFlags err:",err)
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("new config.conf err:", err)
		return
	}

	gvr := schema.GroupVersionResource{
		Group:    "redis.sensoro.sre",
		Version:  "v1beta23",
		Resource: "redis",
	}

	list,err := client.Resource(gvr).Namespace("").List(context.Background(),v1.ListOptions{
		//LabelSelector: fmt.Sprintf("ip=%s",ip),
	})
	if err !=nil{
		fmt.Println("list err:",err)
	}

	for _,v := range list.Items{
		fmt.Println(v.GetName())
	}
}
