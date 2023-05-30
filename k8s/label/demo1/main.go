package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)


var devopsProject = schema.GroupVersionResource{
	Group:    "devops.kubesphere.io",
	Version:  "v1alpha3",
	Resource: "devopsprojects",
}

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("1111:", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("222:", err)
	}

	workspaceLabel := "kubesphere.io/workspace"
	workspaceName := "devops-workspace"

	obj, err := dynamicClient.Resource(devopsProject).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s",workspaceLabel,workspaceName),
	})
	if err != nil {
		fmt.Println("333:", err)
	}

	if len(obj.Items) == 0{
		fmt.Println("this item length is 0")
	}

	for i := range obj.Items {
		//labels := item.GetLabels()
		item := obj.Items[i]
		ggn := item.GetGenerateName()
		fmt.Println(ggn)
		/*
		判断是否有相同的devops-project名字：
		1. jenkins请求接口： http://xxx:30180/view/all/checkJobName?value=devops
		2. 我们需要判断在同一个workspace下是否有相同的name的，可以根据label获取到同一个workspace的所有list，然后再判断name是否有相同的

		 */


		//fmt.Printf("Resource: %s\n", item.GetName())
	}

}
