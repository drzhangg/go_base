package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"os"
	"path/filepath"
)

func main() {
	// 创建 In-Cluster 配置
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 创建 Metrics 客户端
	metricsClient, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 获取所有节点的信息
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 获取节点的资源使用情况
	nodeMetricsList, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 打印表头
	fmt.Printf("%-20s %-15s %-10s %-15s %-15s %-15s %-10s %-10s\n", "节点", "CPU用量", "CPU总量", "CPU平均负载", "内存用量", "内存总量", "磁盘用量", "容器组用量")

	// 计算并打印每个节点的资源用量
	for _, node := range nodes.Items {
		nodeName := node.Name
		cpuCapacity := node.Status.Capacity.Cpu().MilliValue()
		memoryCapacity := node.Status.Capacity.Memory().Value()
		podCapacity := node.Status.Capacity.Pods().Value()
		diskCapacity := node.Status.Capacity.StorageEphemeral().Value() // 假设磁盘使用 StorageEphemeral 统计

		// 获取节点的使用量
		var cpuUsage, memoryUsage int64
		for _, nodeMetrics := range nodeMetricsList.Items {
			if nodeMetrics.Name == nodeName {
				cpuUsage = nodeMetrics.Usage.Cpu().MilliValue()
				memoryUsage = nodeMetrics.Usage.Memory().Value()
			}
		}

		// 获取节点的平均负载（此处假设 Metrics Server 提供的平均负载是最近1分钟的平均值）
		var cpuLoad float64
		loadMetrics, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
		if err == nil {
			load1, ok := loadMetrics.Status.Allocatable["load1"]
			if ok {
				cpuLoad = load1.AsApproximateFloat64()
			}
		}

		// 打印每个节点的资源用量信息
		fmt.Printf("%-20s %-15d %-10d %-15.2f %-15d %-15d %-15d %-10d\n",
			nodeName,
			cpuUsage,
			cpuCapacity,
			float64(cpuUsage)/float64(cpuCapacity)*100, // 计算平均负载（占比）
			memoryUsage/1024/1024,                      // 转换为Mi
			memoryCapacity/1024/1024,                   // 转换为Mi
			diskCapacity/1024/1024,                     // 转换为Mi
			podCapacity,
		)
	}
}
