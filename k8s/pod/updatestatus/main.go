package main

import (
	"context"
	"flag"
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	h := os.Getenv("HOME")

	f := filepath.Join(h, ".kube", "config.conf")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	pod, err := client.CoreV1().Pods("default").Get(context.Background(), "endpoint-pod", v1.GetOptions{})
	if err != nil {
		fmt.Println("get pod err:", err)
	}

	//dataMap := map[string]interface{}{
	//	"status": map[string]interface{}{
	//		"message": "test message",
	//	},
	//}


	//data, err := json.Marshal(&dataMap)
	//if err != nil {
	//	fmt.Println("marshal err:", err)
	//}

	status := pod.Status
	status.Message = "test message"

	_, err = client.CoreV1().Pods(pod.Namespace).
		UpdateStatus(context.Background(), &v12.Pod{
			TypeMeta:   pod.TypeMeta,
			ObjectMeta: pod.ObjectMeta,
			Spec:       pod.Spec,
			Status:     status,
		}, v1.UpdateOptions{})
	//Patch(context.Background(), pod.Name, types.MergePatchType, data, v1.PatchOptions{})
	if err != nil {
		fmt.Println("patch err:", err)
	}

}
