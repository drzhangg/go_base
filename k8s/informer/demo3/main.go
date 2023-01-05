package main

import (
	"fmt"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/signal"
)

func mustClientSet() kubernetes.Interface {
	kubeconfig := os.Getenv("KUBECONFIG")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	return clientset
}

func newConfigMapsListerWatcher() cache.ListerWatcher {
	clientset := mustClientSet()
	// 客户端请求器
	client := clientset.CoreV1().RESTClient()
	resource := "configmaps"
	namespace := "tmp"

	selector := fields.Everything()
	lw := cache.NewListWatchFromClient(client, resource, namespace, selector)
	return lw
}

func main() {
	fmt.Println("----- 1-list-watcher -----")

	lw := newConfigMapsListerWatcher()

	list, err := lw.List(v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	items, err := meta.ExtractList(list)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Initial list:")

	for _, item := range items {
		configMap, ok := item.(*v12.ConfigMap)
		if !ok {
			return
		}
		fmt.Println(configMap.Name)

		accessor, err := meta.Accessor(item)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(accessor.GetName())
	}

	listMetaInterface, err := meta.ListAccessor(list)
	if err != nil {
		fmt.Println(err)
	}

	// resourceVersion在同步的过程中非常重要
	resourceVersion := listMetaInterface.GetResourceVersion()

	w, err := lw.Watch(v1.ListOptions{
		ResourceVersion: resourceVersion,
	})
	if err != nil {
		fmt.Println(err)
	}

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, os.Interrupt)
	fmt.Println("start watching...")

loop:
	for {
		select {
		case <-stopCh:
			fmt.Println("Interrupted")
			break loop
		case event, ok := <-w.ResultChan():
			if !ok {
				fmt.Println("Broken channel")
				break loop
			}
			configMap, ok := event.Object.(*v12.ConfigMap)
			if !ok {
				return
			}
			fmt.Printf("%s: %s\n", event.Type, configMap.Name)
		}
	}
}
