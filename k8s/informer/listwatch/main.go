package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func client() *kubernetes.Clientset {
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

	return clientset
}

func main() {
	client := client()

	podLW := cache.NewListWatchFromClient(client.CoreV1().RESTClient(),"pods","default",fields.Everything())


	watch ,err := podLW.Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case v, ok := <-watch.ResultChan():
			if ok {
				fmt.Println(v.Type, ":",v.Object.(*v1.Pod).Name)
			}
		}
	}


	//list,err := podLW.List(metav1.ListOptions{})
	//if err!= nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Printf("%T\n",list)
	//
	//podList := list.(*v1.PodList)
	//
	//for _, item := range podList.Items{
	//
	//	fmt.Println(item.Name)
	//}

}
