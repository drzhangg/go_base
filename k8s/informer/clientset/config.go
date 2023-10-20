package clientset

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

const kubeConfigFilePath = "/Users/drzhang/.kube/config"


var (
	Resource = "pods"
	Namespace = "default"
)

type K8sConfig struct {
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

// 读取kubeconfig 配置文件
func (this *K8sConfig) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigFilePath)

	if err != nil {
		log.Fatal(err)
	}

	return config
}

// 初始化 clientSet
func (this *K8sConfig) InitClient() *kubernetes.Clientset {
	c, err := kubernetes.NewForConfig(this.K8sRestConfig())

	if err != nil {
		log.Fatal(err)
	}

	return c
}

// 初始化 dynamicClient
func (this *K8sConfig) InitDynamicClient() dynamic.Interface {
	c, err := dynamic.NewForConfig(this.K8sRestConfig())

	if err != nil {
		log.Fatal(err)
	}

	return c
}

// 初始化 DiscoveryClient
func (this *K8sConfig) InitDiscoveryClient() *discovery.DiscoveryClient {
	return discovery.NewDiscoveryClient(this.InitClient().RESTClient())
}
