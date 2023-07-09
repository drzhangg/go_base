package main

import (
	"fmt"
	"go_base/k8s/informer/base"
	"go_base/k8s/informer/sample/dog"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

type PodHandler struct {
}

func (p PodHandler) OnAdd(obj interface{}) {
	//TODO implement me
	fmt.Println("OnAdd :", obj.(*v1.Pod).Name)
}

func (p PodHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
	fmt.Println("OnUpdate :", newObj.(*v1.Pod).Name)
}

func (p PodHandler) OnDelete(obj interface{}) {
	//TODO implement me
	fmt.Println("OnDelete :", obj.(*v1.Pod).Name)
}

var _ cache.ResourceEventHandler = &PodHandler{}

func main() {
	client := base.KubeClient()
	podLw := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "pods", "default", fields.Everything())

	wd := dog.NewWatchDog(podLw, &v1.Pod{}, PodHandler{})
	wd.Run()
}
