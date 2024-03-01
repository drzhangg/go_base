package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
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
		fmt.Println("rest.InClusterConfig err:",err)
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}
	// 创建 Clientset 对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	dc, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("new config.conf err:", err)
		return
	}

	result,err := dc.Resource(redisClusterGVR).List(context.Background(),v1.ListOptions{})
	if err != nil {
		fmt.Println("gvr List err:", err)
		return
	}

	data,_ := result.MarshalJSON()

	fmt.Println("data:",string(data))

	rc := clientset.RESTClient()

	r := rc.Get().AbsPath("/apis/redis.sensoro.sre/v1beta2/redisclusters")

	data1, _ := r.DoRaw(context.Background())

	fmt.Println("data1:",string(data1))


}

var (
	redisClusterGVR = schema.GroupVersionResource{
		Group:    "redis.sensoro.sre",
		Version:  "v1beta2",
		Resource: "redisclusters",
	}
)
