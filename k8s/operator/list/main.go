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

	f := filepath.Join(h, ".kube", "config.conf")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {

	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("new config.conf err:", err)
		return
	}

	gvr := schema.GroupVersionResource{
		Group:    "app.drzhangg.io",
		Version:  "v1beta1",
		Resource: "appservices",
	}

	un,err := client.Resource(gvr).List(context.Background(),v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	if len(un.Items) == 0{
		fmt.Println("this item length is 0")
	}

	//result := []string{}
	for _,item := range un.Items {
		_, ok := item.GetLabels()["ip"]
		if !ok {
			continue
		}

		fmt.Println("ip::", item.GetObjectKind().GroupVersionKind().Kind)
		fmt.Println("name::", item.GetName())
	}
}
