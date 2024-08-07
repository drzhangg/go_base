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

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
	namespace := "sre"
	podName := "wakanda-6bdb64767d-j6sgg" // 替换为你的 Pod 名称

	podLogOptions := corev1.PodLogOptions{
		Follow:    true, // 如果需要实时跟踪日志，请将这个设置为 true
		Container: "wakanda",
	}

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade to websocket: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// 创建日志请求
		request := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)

		podLogs, err := request.Stream(context.Background())
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Failed to get logs: "+err.Error()))
			return
		}
		defer podLogs.Close()

		buf := make([]byte, 1024)
		for {
			n, err := podLogs.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				conn.WriteMessage(websocket.TextMessage, []byte("Error reading logs: "+err.Error()))
				return
			}
			if n > 0 {
				conn.WriteMessage(websocket.TextMessage, buf[:n])
			}
		}
	})

	// 启动 HTTP 服务器
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
