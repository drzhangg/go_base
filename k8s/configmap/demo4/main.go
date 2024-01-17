package main

import (
	"context"
	"encoding/json"
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

	cm,err := clientset.CoreV1().ConfigMaps("kube-system").Get(context.TODO(),"coredns",v1.GetOptions{})
	if err != nil {
		fmt.Println("err:", err)
	}

	data,err := json.Marshal(cm)
	if err != nil {
		fmt.Println("err23:",err)
	}

	fmt.Printf("%s",data)


}
