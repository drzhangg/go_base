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

func formatSize(size int64) string {
	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
	)

	switch {
	case size >= TiB:
		return fmt.Sprintf("%.2fTi", float64(size)/TiB)
	case size >= GiB:
		return fmt.Sprintf("%.2fGi", float64(size)/GiB)
	case size >= MiB:
		return fmt.Sprintf("%.2fMi", float64(size)/MiB)
	default:
		return fmt.Sprintf("%dKi", size/KiB)
	}
}

func formatCPU(size int64) string {
	const (
		mCPU = 1000
	)

	if size >= mCPU {
		return fmt.Sprintf("%.2f", float64(size)/mCPU)
	}
	return fmt.Sprintf("%dm", size)
}

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	metricsClient, err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Println("metrics error:", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("clientset error:", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Println("list node err:", err)
		return
	}

	nodeMetricsList, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Println("metrics node err:", err)
		return
	}

	// 用于统计的变量
	nodeUsageMap := make(map[string]map[string]int64)

	// 计算总量
	for _, node := range nodes.Items {
		nodeName := node.Name
		cpu := node.Status.Capacity.Cpu().MilliValue()
		memory := node.Status.Capacity.Memory().Value()
		pods := node.Status.Capacity.Pods().Value()
		// Kubernetes nodes don't have StorageEphemeral directly in their capacity
		// disk := node.Status.Capacity.StorageEphemeral().Value() // 假设磁盘使用 StorageEphemeral 统计

		nodeUsageMap[nodeName] = map[string]int64{
			"totalCPU":    cpu,
			"totalMemory": memory,
			"totalPods":   pods,
			// "totalDisk":   disk,
		}
	}

	// 计算使用量
	for _, nodeMetrics := range nodeMetricsList.Items {
		nodeName := nodeMetrics.Name
		usedCPU := nodeMetrics.Usage.Cpu().MilliValue()
		usedMemory := nodeMetrics.Usage.Memory().Value()

		nodeUsageMap[nodeName]["usedCPU"] = usedCPU
		nodeUsageMap[nodeName]["usedMemory"] = usedMemory

		// nodeMetrics.Usage.StorageEphemeral().Value() // 假设磁盘使用 StorageEphemeral 统计
	}

	// 输出统计结果
	for nodeName, usage := range nodeUsageMap {
		totalCPU := usage["totalCPU"]
		usedCPU := usage["usedCPU"]
		totalMemory := usage["totalMemory"]
		usedMemory := usage["usedMemory"]
		totalPods := usage["totalPods"]
		// totalDisk := usage["totalDisk"]

		cpuUsagePercent := float64(usedCPU) / float64(totalCPU) * 100
		memoryUsagePercent := float64(usedMemory) / float64(totalMemory) * 100

		fmt.Printf("Node: %s\n", nodeName)
		fmt.Printf("  Total CPU: %s\n", formatCPU(totalCPU))
		fmt.Printf("  Used CPU: %s (%.2f%%)\n", formatCPU(usedCPU), cpuUsagePercent)
		fmt.Printf("  Total Memory: %s\n", formatSize(totalMemory))
		fmt.Printf("  Used Memory: %s (%.2f%%)\n", formatSize(usedMemory), memoryUsagePercent)
		fmt.Printf("  Total Pods: %d\n", totalPods)
		// fmt.Printf("  Total Disk: %s\n", formatSize(totalDisk))
	}
}
