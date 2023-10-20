package main

import (
	"fmt"
	clientset2 "go_base/k8s/informer/clientset"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"time"
)

// simulate  K8s simple reflector creation process
func main() {

	config := clientset2.K8sConfig{}

	cliset := config.InitClient()


	//cliset := helper.InitK8SClient()
	// 使用 store 进行存储，这样本地才有一份数据；
	// 如果本地没有存储到被删除的资源， 则不需要 Pop 该资源的 Delete 事件；
	// 所以我们为了准确接收到delete时接收到 Delete 事件, 所以预先创建一下 store
	// cache.MetaNamespaceKeyFunc 是用于返回资源的唯一标识, {namespace}/{name} 或 {name}
	store := cache.NewStore(cache.MetaNamespaceKeyFunc)

	// create list & watch Client
	lwc := cache.NewListWatchFromClient(cliset.CoreV1().RESTClient(),
		clientset2.Resource,
		clientset2.Namespace,
		fields.Everything(),
	)

	// create deltafifo
	df := cache.NewDeltaFIFOWithOptions(
		cache.DeltaFIFOOptions{
			KeyFunction:  cache.MetaNamespaceKeyFunc,
			KnownObjects: store,
		})

	// crete reflector
	rf := cache.NewReflector(lwc, &v1.Pod{}, df, time.Second*0)
	rsCH := make(chan struct{})
	go func() {
		rf.Run(rsCH)
	}()

	// fetch delta event
	for {
		df.Pop(func(i interface{}) error {
			// deltas
			for _, d := range i.(cache.Deltas) {
				fmt.Println(d.Type, ":", d.Object.(*v1.Pod).Name,
					"-", d.Object.(*v1.Pod).Status.Phase)
				switch d.Type {
				case cache.Sync, cache.Added:
					// 向store中添加对象
					store.Add(d.Object)
				case cache.Updated:
					store.Update(d.Object)
				case cache.Deleted:
					store.Delete(d.Object)
				}
			}
			return nil
		})
	}
}