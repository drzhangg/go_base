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

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}

	client,err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("get client err:",err)
	}

	pod,err := client.CoreV1().Pods("default").Get(context.Background(),"test",v1.GetOptions{})
	if err != nil {
		fmt.Println("get pod err:",err)
		return
	}

	labels := pod.Labels
	if labels == nil {
		labels = make(map[string]string)
	}

	if _,ok :=labels["app"];!ok{
		labels["app"] = "test"
	}
	pod.Labels = labels

	client.CoreV1().Pods("default").Update(context.Background(),pod,v1.UpdateOptions{})
}
