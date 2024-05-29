package main

import (
	"context"
	"flag"
	"fmt"
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

	// HTTP 处理函数
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade to websocket: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// 获取 namespace 和 pod name
		namespace := r.URL.Query().Get("namespace")
		podName := r.URL.Query().Get("podName")
		containerName := r.URL.Query().Get("container")

		// 设置日志选项
		podLogOptions := corev1.PodLogOptions{
			Follow:    true,
			Container: containerName,
		}

		// 创建日志请求
		request := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)
		podLogs, err := request.Stream(context.Background())
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Failed to get logs: "+err.Error()))
			return
		}
		defer podLogs.Close()

		// 使用 Channel 传输日志数据
		logChan := make(chan []byte)
		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := podLogs.Read(buf)
				if err != nil {
					if err == io.EOF {
						close(logChan)
						break
					}
					logChan <- []byte("Error reading logs: " + err.Error())
					close(logChan)
					break
				}
				if n > 0 {
					logChan <- buf[:n]
				}
			}
		}()

		// 通过 WebSocket 实时输出日志数据
		for msg := range logChan {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("write message error: ", err)
				break
			}
		}
	})

	// 启动 HTTP 服务器
	if err := http.ListenAndServe(":8084", nil); err != nil {
		panic(err)
	}
}
