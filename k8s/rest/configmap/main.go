package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	rc := clientset.RESTClient()

	// api/v1/namespaces/sre/configmaps
	//r := rc.Get().AbsPath("/apis/redis.sensoro.sre/v1beta2/redisclusters")
	r := rc.Get().AbsPath("api/v1/namespaces/sre/configmaps")

	data1, _ := r.DoRaw(context.Background())

	//fmt.Println("data1:",string(data1))

	body := struct {
		v1.ListMeta `json:"metadata"`
		Items        []v12.ConfigMap `json:"items"`
	}{}

	json.Unmarshal(data1,&body)

	//fmt.Println(body.Items)

	for _,val := range body.Items{
		fmt.Println("kind:",val.Kind)
	}



}
