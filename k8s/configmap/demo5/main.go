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
	name := "wakanada-role-config"


	cm ,err := clientset.CoreV1().ConfigMaps(ns).Get(context.Background(),name,v1.GetOptions{})
	if err != nil {
		fmt.Println("cm err:",err)
	}

	//fmt.Println(cm.Data)

	//user := cm.Data["user"]

	//m := make(map[string]map[string][]string)
	for k, _ := range cm.Data{
		fmt.Println("k::",k)

		fmt.Println("k1::",cm.Data[k])


		//for _,val1 := range cm.Data[k]{
		//	fmt.Println("val::",string(val1))
		//}

		//for _, val1 := range cm.Data[val]{
		//	arr1 := strings.Split(cm.Data[val][val1],":")
		//	verbs := strings.Split(arr1[1],",")
		//}
	}

	//fmt.Println(user)

	//fmt.Println(reflect.TypeOf(user))


	//
	//fmt.Println(verbs)
	//
	//
	//m[arr1[0]] = append([]string{}, verbs...)
	//
	//fmt.Println(m)
}
