package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func mustClientSet() kubernetes.Interface {
	kubeconfig := os.Getenv("KUBECONFIG")

	config,err := clientcmd.BuildConfigFromFlags("",kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset,err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func main() {
	clientset := mustClientSet()

	ns,err := clientset.CoreV1().Namespaces().List(context.TODO(),v1.ListOptions{})
	if err !=nil{
		panic(err)
	}

	for _, n := range ns.Items{
		fmt.Println(n.Name)
	}
}
