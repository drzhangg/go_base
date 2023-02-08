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

// 代码从configmap中读取值
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

	cms, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println("err:", err)
	}

	gvkMap := map[string]GVR{}

	for _, v := range cms.Items {
		if v.Name == "test-cm" {
			//fmt.Println(v.Data)
			fmt.Println("len::", len(v.Data))
			for k, v := range v.Data {
				//fmt.Println("key:", k, "val:", v)

				var d GVR
				err := yaml.Unmarshal([]byte(v),&d)
				if err !=nil{
					fmt.Println(err)
				}
				gvkMap[k] = d
			}
		}
	}
	fmt.Println(gvkMap)
}

type GVR struct {
	Version string `json:"Version"`
	Kind    string `json:"Resource"`
	Group   string `json:"Group"`
}

