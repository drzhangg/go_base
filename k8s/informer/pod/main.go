package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
)

func main() {
	h := os.Getenv("HOME")

	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		fmt.Println("err:",err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("err1:",err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	t := time.Now()


	sharedInformers := informers.NewSharedInformerFactory(clientset, 0)
	podInformer := sharedInformers.Core().V1().Pods()

	informer := podInformer.Informer()

	podLister := podInformer.Lister()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})

	sharedInformers.Start(stopCh)
	sharedInformers.WaitForCacheSync(stopCh)

	pods,err := podLister.Pods("").List(labels.Everything())
	if err != nil {
		fmt.Println("err3:",err)
	}

	//pods = pods

	for _, val := range pods{
		fmt.Println("nmae::",val.Name)
		if val.GetName() == "mongodb-2" {
			fmt.Println("val::::",val)
		}
	}

	fmt.Println("time::",time.Now().Sub(t))

	<-stopCh

}

func onAdd(obj interface{}) {
	//pod := obj.(*v1.Pod)
	//fmt.Println("add a pod:", pod.Name)
}

func onUpdate(old, new interface{}) {
	//oldDeploy := old.(*v1.Pod)
	//newDeploy := new.(*v1.Pod)
	//fmt.Println("update pod:", oldDeploy.Name, newDeploy.Name)
}

func onDelete(obj interface{}) {
	//pod := obj.(*v1.Pod)
	//fmt.Println("delete a pod:", pod.Name)
}
