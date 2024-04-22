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

	ns := "sre"
	name := "dbaudit-archery-d7644ff4-9b9mf"
	event,err := clientset.CoreV1().Events(ns).List(context.Background(),v1.ListOptions{})
	if err != nil {
		fmt.Println("get event err:",err)
		return
	}

	for _,val := range event.Items{

		if val.InvolvedObject.Kind != "Pod" || val.InvolvedObject.Name != name{
			continue
		}

		//fmt.Printf("last:%#v\n",val.LastTimestamp)
		//fmt.Printf("type:%#v\n",val.Type)
		//fmt.Printf("reason:%#v\n",val.Reason)
		//fmt.Printf("name:%#v\n",val.ObjectMeta.Name)
		fmt.Printf("source:%#v\n",val.Source)
		//fmt.Printf("Related:%#v\n",val.InvolvedObject)


		//fmt.Printf("message:%#v\n",val.Message)


	}


}
