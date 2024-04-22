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

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	ns := "default"
	name:= "test1"

	secret ,err := clientset.CoreV1().Secrets(ns).Get(context.Background(),name,v1.GetOptions{})
	if err != nil {
		fmt.Println("get secret err:",err)
	}

	fmt.Println("secret::",secret.Data)
	fmt.Println("secret1::",string(secret.Data["user_pass"]))

}
