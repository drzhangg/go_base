package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
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

	ns := "sre"
	name := "wakanda-role-config"

	cm ,err := clientset.CoreV1().ConfigMaps(ns).Get(context.Background(),name,v1.GetOptions{})
	if err != nil {
		fmt.Println("cm err:",err)
	}


	conf := cm.Data["role.conf"]
	fmt.Println("role::",conf)
	var auth Roles

	err = yaml.Unmarshal([]byte(conf),&auth)
	if err != nil {
		fmt.Println("unmarsharl err:",err)
	}

	fmt.Println("auth::",auth)
}

type Roles struct {
	Authorities []Authority
}

type Authority struct {
	Role string
	Rules []Rule
}

type Rule struct {
	Resource string
	Verbs map[string]bool
}
