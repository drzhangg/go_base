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

	deploys,err := clientset.AppsV1().Deployments("").List(context.TODO(),v1.ListOptions{
		//LabelSelector: "control-plane",
	})
	if err != nil {
		fmt.Println("get deploy list err:", err)
	}

	fmt.Printf("one:%#v\n",deploys.Items[0])

	//for _,v := range deploys.Items{
	//	fmt.Println(v.Name)
	//}

}
