package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var vm = schema.GroupVersionResource{
	Group:    "drzhangg",
	Version:  "v1",
	Resource: "apps",
}

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	discoveryClient := clientset.Discovery()

	_,apiresources,err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		fmt.Println("get apiresources failed:",err)
	}

	for _, apiresource := range apiresources{
		for _,v := range apiresource.APIResources{
			fmt.Println(v.Name)
		}
	}



}
