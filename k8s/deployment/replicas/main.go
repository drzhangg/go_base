package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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

	rc := clientset.RESTClient()



	namespace := "default"
	deployName := "test"
	deploy,err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(),deployName,v1.GetOptions{})
	if err != nil {
		fmt.Println("get deploy err:", err)
	}

	*deploy.Spec.Replicas = 4

	data,_ := deploy.Marshal()

	deployPath := "apis/apps/v1/namespaces/default/deployments"

	r := rc.Patch(types.StrategicMergePatchType).AbsPath(deployPath).Body(data)

	data1, err := r.DoRaw(context.Background())
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("dd:",string(data1))


	//
	//_,err = clientset.AppsV1().Deployments(namespace).Update(context.Background(),deploy,v1.UpdateOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	fmt.Println("Deployment 扩容成功")



}
