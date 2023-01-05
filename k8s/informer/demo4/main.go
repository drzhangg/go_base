package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"
)

// newStore creates a new cache.Store.
func newStore() cache.Store {
	return cache.NewStore(cache.MetaNamespaceKeyFunc)
}

// newQueue 用于创建一个 cache.Queue 对象，这里实现为 FIFO 先进先出队列
func newQueue(store cache.Store) cache.Queue {
	return cache.NewDeltaFIFOWithOptions(cache.DeltaFIFOOptions{
		KnownObjects:          store,
		EmitDeltaTypeReplaced: true,
	})
}

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

func newConfigMapsReflector(queue cache.Queue) *cache.Reflector {
	lw := newConfigMapsListerWatcher()
	return cache.NewReflector(lw, &v1.ConfigMap{}, queue, 0)
}

func main() {
	fmt.Println("----- 2-reflector -----")

	store := newStore()
	queue := newQueue(store)
	reflector := newConfigMapsReflector(queue)

	stopCh := make(chan struct{})
	defer close(stopCh)

	// reflector 开始运行后，队列中就会推入新收到的事件
	go reflector.Run(stopCh)


	processOjb := func(obj interface{}) error {
		for _, d:= range obj.(cache.Deltas){
			switch d.Type {
			case cache.Sync, cache.Added, cache.Updated,cache.Replaced:
				if _,exists,err := store.Get(d.Object); err == nil && exists {
					if err =store.Update(d.Object);err != nil {
						return err
					}
				} else {
					if err = store.Add(d.Object);err != nil {
						return err
					}
				}
			case cache.Deleted:
				if err := store.Delete(d.Object);err !=nil{
					return err
				}
			}

			configMap,ok := d.Object.(*v1.ConfigMap)
			if !ok {
				return fmt.Errorf("not config: %T", d.Object)
			}
			fmt.Printf("%s: %s\n", d.Type, configMap.Name)
		}
		return nil
	}

	fmt.Println("Start syncing...")

	wait.Until(func() {
		for {
			_, err := queue.Pop(processOjb)
			if err != nil {
				fmt.Println(err)
			}
		}
	}, time.Second, stopCh)

}
