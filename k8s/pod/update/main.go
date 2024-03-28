package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	rc := client.RESTClient()

	deploy, err := client.AppsV1().Deployments("default").Get(context.Background(), "test1", metav1.GetOptions{})
	if err != nil {
		fmt.Println("get deployment err:", err)
	}

	*deploy.Spec.Replicas = 4

	//deploy.Kind = "Deployment"
	//deploy.APIVersion = "apps/v1"

	deployPath := "apis/apps/v1/namespaces/default/deployments/test1"
	fmt.Println(deployPath)
	//data,err := deploy.Marshal()

	data,err := json.Marshal(deploy)
	if err!= nil {
		fmt.Println("marshal err:",err)
	}
	fmt.Println("kind:",string(data))

	result,err := rc.Put().AbsPath(deployPath).Body(data).DoRaw(context.Background())
	if err != nil {
		fmt.Println("pu err::",err)
	}

	//re,err := rc.Put().
	//	AbsPath(deployPath).
	//	Namespace("default").
	//	Resource("deployments").
	//	Name("test1").
	//	VersionedParams(&metav1.UpdateOptions{}, scheme.ParameterCodec).
	//	Body(data).
	//	DoRaw(context.Background())
	//	//Into(result)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}


	fmt.Println("result:",string(result))


}
