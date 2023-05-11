package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"time"
)

func main() {
	var err error
	var config *rest.Config

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "[可选] kubeconfig 绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 绝对路径")
	}
	// 初始化 rest.Config 对象
	if config, err = rest.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}
	// 创建 Clientset 对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 初始化informer factory（为了测试方便这里设置每30s重新 List 一次）
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)

	// 对 deployment 监听
	deployInformer := informerFactory.Apps().V1().Deployments()

	podInformer := informerFactory.Core().V1().Pods()

	// 创建Informer（相当于注册到工厂中，这样下面启动的时候就会List&Watch对应的资源）
	informer := deployInformer.Informer()

	podinfor := podInformer.Informer()

	// 创建Lister
	deployLister := deployInformer.Lister()
	podLister := podInformer.Lister()

	podinfor.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: onAdd1,
		UpdateFunc: onUpdate1,
		DeleteFunc: onDelete1})

	// 注册事件处理程序
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})

	stopper := make(chan struct{})
	defer close(stopper)

	// 启动 informer，list & watch
	informerFactory.Start(stopper)
	// 等待所有启动的 Informer 的缓存被同步
	informerFactory.WaitForCacheSync(stopper)

	// 从本地缓存中获取default中的所有deployment列表
	ds, err := deployLister.Deployments("default").List(labels.Everything())
	if err != nil {
		panic(err)
	}

	for k, v := range ds {
		fmt.Printf("%d -> %s\n", k+1, v.Name)
	}

	podlist, err := podLister.Pods("default").List(labels.Everything())
	if err != nil {
		panic(err)
	}

	for _, item := range podlist {
		fmt.Println(item.Name)
	}

	<-stopper
}
func onAdd1(obj interface{}) {
	deploy := obj.(*v12.Pod)
	fmt.Println("add a Pod:", deploy.Name)
}

func onUpdate1(old, new interface{}) {
	oldDeploy := old.(*v12.Pod)
	newDeploy := new.(*v12.Pod)
	fmt.Println("update Pod:", oldDeploy.Name, newDeploy.Name)
}

func onDelete1(obj interface{}) {
	deploy := obj.(*v12.Pod)
	fmt.Println("delete a Pod:", deploy.Name)
}

func onAdd(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	fmt.Println("add a deployment:", deploy.Name)
}

func onUpdate(old, new interface{}) {
	oldDeploy := old.(*v1.Deployment)
	newDeploy := new.(*v1.Deployment)
	fmt.Println("update deployment:", oldDeploy.Name, newDeploy.Name)
}

func onDelete(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	fmt.Println("delete a deployment:", deploy.Name)
}
