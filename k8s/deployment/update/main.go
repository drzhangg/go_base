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

	deploy, err := clientset.AppsV1().Deployments("default").Get(context.Background(), "test-nginx", v1.GetOptions{})
	if err != nil {
		fmt.Println("get deploy err:", err)
	}

	deploy.Spec.Template.Spec.Containers[0].Image = "nginx:1.7"
	var replicas int32 = 2
	rs := &replicas
	deploy.Spec.Replicas = rs

	deploy.Status.Conditions[0].Message = "test message"

	_, err = clientset.AppsV1().Deployments("default").Update(context.TODO(), deploy, v1.UpdateOptions{})
	if err != nil {
		fmt.Println(err)
	}

}
