package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
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

	cm, err := clientset.CoreV1().ConfigMaps("default").Get(context.TODO(), "cm-demo", v1.GetOptions{})
	if err != nil {
		fmt.Println("get cm err:", err)
		if errors.IsNotFound(err){
			fmt.Println("the cm is not found!")
			return
		}
	}

	var m = make(map[string]int)
	//fmt.Println(cm.Data)
	for key, _ := range cm.Data {
		//fmt.Println("key:",key)
		m[key] = 0
	}
	fmt.Println(m)

}
