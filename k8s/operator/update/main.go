package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
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

	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("new config.conf err:", err)
		return
	}

	gvr := schema.GroupVersionResource{
		Group:    "data.my.domain",
		Version:  "v1beta1",
		Resource: "redis",
	}

	redisCr,err := client.Resource(gvr).Get(context.Background(),"redis-sample",v1.GetOptions{})
	if err != nil {
		fmt.Println("get redis crd err:", err)
		return
	}

	m := redisCr.GetAnnotations()
	if _,ok := m["random"];!ok {
		m["random"] = "asdasdqwewq"
	}
	redisCr.SetAnnotations(m)

	//if redisCr.GetAnnotations()["kubectl.kubernetes.io/last-applied-configuration"] == ""{
	//	redisCr.SetAnnotations(annotation)
	//}

	_,err  = client.Resource(gvr).Update(context.Background(),redisCr,v1.UpdateOptions{})
	if err != nil {
		fmt.Println("updatestatus redis crd err:", err)
		return
	}


	redisList,err := client.Resource(gvr).List(context.Background(),v1.ListOptions{})
	if err != nil {
		fmt.Println("list redis err:", err)
		return
	}

	for _, v := range redisList.Items{
		fmt.Println("name:",v.GetName())
	}


}
