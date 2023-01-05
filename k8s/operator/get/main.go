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

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {

	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("new config err:", err)
		return
	}

	gvr := schema.GroupVersionResource{
		Group:    "app.drzhangg.io",
		Version:  "v1beta1",
		Resource: "appservices",
	}

	ip := "9a91fb901"


	list,err := client.Resource(gvr).List(context.Background(),v1.ListOptions{
		LabelSelector: fmt.Sprintf("ip=%s",ip),
	})
	if err !=nil{
		fmt.Println("list err:",err)
	}

	for _,v := range list.Items{
		fmt.Println(v.GetName())
	}
}
