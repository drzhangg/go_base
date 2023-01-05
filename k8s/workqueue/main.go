package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

type Controller struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller
}

func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller {
	return &Controller{
		indexer:  indexer,
		queue:    queue,
		informer: informer,
	}
}

func (c *Controller) processNextItem() bool {

	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)



	return false
}

// syncToStdout 是控制器的业务逻辑实现
// 在此控制器中，它只是将有关Pod的信息打印到 stdout
// 如果发生错误，则简单的返回错误
// 此外重试逻辑不应成为业务逻辑的一部分
func (c *Controller) syncToStdout(key string) error {
	//从本地存储中获取key对应的对象
	obj, exists, err := c.indexer.Get(key)
	if err != nil {
		klog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		fmt.Printf("Pod %s does not exists anymore\n", key)
	} else {
		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
	}
	return nil
}

func (c *Controller) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	if c.queue.NumRequeues(key) < 5 {
		klog.Infof("Error syncing pod %v: %v", key, err)

		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)

	runtime.HandleError(err)
	klog.Infof("Dropping pod %q out of the queue: %v", key, err)
}

func main() {

}
