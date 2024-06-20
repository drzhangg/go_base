package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"os"
	"path/filepath"
)

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	metricsClient,err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Println("metrics error:", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("clientset error:", err)
	}

	nodes,err := clientset.CoreV1().Nodes().List(context.Background(),v1.ListOptions{})
	if err != nil{
		fmt.Println("list node err:", err)
		return
	}

	nodeMetricsList,err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(),v1.ListOptions{})
	if err != nil {
		fmt.Println("metrics node err:",err)
		return
	}

	// 用于统计的变量
	var totalCPU, totalMemory, totalPods, totalDisk int64
	var usedCPU, usedMemory int64

	// 计算总量
	for _, node := range nodes.Items {
		cpu := node.Status.Capacity.Cpu().MilliValue()
		memory := node.Status.Capacity.Memory().Value()
		pods := node.Status.Capacity.Pods().Value()
		disk := node.Status.Capacity.StorageEphemeral().Value() // 假设磁盘使用 StorageEphemeral 统计

		totalCPU += cpu
		totalMemory += memory
		totalPods += pods
		totalDisk += disk
	}

	// 计算使用量
	for _, nodeMetrics := range nodeMetricsList.Items {
		usedCPU += nodeMetrics.Usage.Cpu().MilliValue()
		usedMemory += nodeMetrics.Usage.Memory().Value()
	}

	// 计算使用百分比
	cpuUsagePercent := float64(usedCPU) / float64(totalCPU) * 100
	memoryUsagePercent := float64(usedMemory) / float64(totalMemory) * 100

	// 输出统计结果
	fmt.Printf("Total CPU: %dm\n", totalCPU)
	fmt.Printf("Used CPU: %dm (%.2f%%)\n", usedCPU, cpuUsagePercent)
	fmt.Printf("Total Memory: %dMi\n", totalMemory/1024/1024/1024)
	fmt.Printf("Used Memory: %dMi (%.2f%%)\n", usedMemory/1024/1024/1024, memoryUsagePercent)
	fmt.Printf("Total Pods: %d\n", totalPods)
	fmt.Printf("Total Disk: %dMi\n", totalDisk/1024/1024/1024)



}
