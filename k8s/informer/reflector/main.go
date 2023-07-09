package main

import (
	"fmt"
	"go_base/k8s/informer/base"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

func main() {

	client := base.KubeClient()

	podLw := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "pods", "default", fields.Everything())

	df := cache.NewDeltaFIFOWithOptions(cache.DeltaFIFOOptions{
		KeyFunction: cache.MetaNamespaceKeyFunc,
	})

	rf := cache.NewReflector(podLw, &v1.Pod{}, df, 0)

	ch := make(chan struct{})

	go func() {
		rf.Run(ch)
	}()

	for {
		df.Pop(func(obj interface{}) error {
			for _, delta := range obj.(cache.Deltas) {
				fmt.Println(delta.Type, ":", delta.Object.(*v1.Pod).Name, ":", delta.Object.(*v1.Pod).Status.Phase)
			}

			return nil
		})
	}
}
