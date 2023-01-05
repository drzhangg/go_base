package main

import (
	"context"
	"flag"
	"fmt"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	bindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), v12.ListOptions{})
	if err != nil {
		fmt.Println("get cluster role bindings list err:", err)
		return
	}
	//fmt.Println("binding::",bindings.Items)

	for _, v := range bindings.Items {
		fmt.Println("rolebinding name:", v.Name)
		//fmt.Println("subject::", v.Subjects)

		//fmt.Println("name:",v.Name ," ,Namespace:",v.RoleRef.Name," ,Kind:",v.RoleRef.Kind)
	}



}
