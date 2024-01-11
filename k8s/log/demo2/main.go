package main

import (
	"context"
	"flag"
	"io"
	"net/http"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// 使用 Service Account 或在机器上的 kubeconfig 初始化 client-go 客户端
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(可选) kubeconfig 文件的绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 文件的绝对路径")
	}
	flag.Parse()

	// 使用 kubeconfig 文件构建一个 config 对象
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 获取指定 Pod 的日志
	namespace := "kube-system"
	podName := "coredns-5d78c9869d-2p5rx" // 替换为你的 Pod 名称

	podLogOptions := corev1.PodLogOptions{
		Follow: true, // 如果需要实时跟踪日志，请将这个设置为 true
	}

	request := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)

	podLogs, err := request.Stream(context.Background())
	if err != nil {
		panic(err.Error())
	}
	defer podLogs.Close()

	// 读取和处理日志流
	// 例如，直接打印到标准输出
	//buffer := make([]byte, 2000)
	//for {
	//	numBytes, err := podLogs.Read(buffer)
	//	if numBytes == 0 {
	//		continue
	//	}
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//
	//	msg := string(buffer[:numBytes])
	//	fmt.Print(msg)
	//
	//	// 执行你的逻辑，比如发送到一个 Web 接口等
	//	// ...
	//}
	//
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		// ... set up client-go, get pod name from request, and set up log options

		logStream, err := request.Stream(context.Background())
		if err !=nil {
			return
		}
		// 错误处理省略
		defer logStream.Close()

		// copy the stream to the response writer
		io.Copy(w, logStream)
	})

	http.ListenAndServe(":8080", nil)

}
