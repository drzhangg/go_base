package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	cmList, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println("get cm list err:", err)
	}

	var ma = make(map[string]int)
	for _, v := range cmList.Items {
		if v.Name == "cm-demo" {
			for key, _ := range v.Data {
				ma[key] = 0
			}
			//fmt.Println("data:",v.Data)
		}
	}
	fmt.Println("ma::", ma)

}
