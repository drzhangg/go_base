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

	pvcList,err := clientset.CoreV1().PersistentVolumeClaims("").List(context.TODO(),v1.ListOptions{})
	if err !=nil{
		fmt.Println("err:",err)
	}

	for _,v := range pvcList.Items{
		if v.Spec.StorageClassName != nil {
			fmt.Println(*v.Spec.StorageClassName)
		}else {
			fmt.Println(v.Name)
		}
	}
}
