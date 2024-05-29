package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	var line int64 = 1000

	ns := "sre"
	pod := "wakanda-6bdb64767d-j6sgg"

	res, err := clientset.CoreV1().Pods(ns).
		GetLogs(pod,
			&v1.PodLogOptions{Follow: true, TailLines: &line,Container: "wakanda"}).
		Stream(context.Background())
	if err != nil {
		fmt.Println("get log err:", err)
		return
	}

	reader := bufio.NewReader(res)

	//msgChan := make(chan []byte, 2048)

	//go recieveMessage(reader, msgChan)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("failed to read message: %v\n", err)
			return
		}

		if err == io.EOF {
			break
		}

		fmt.Print("log::", string(bytes))
	}

	//var buf bytes.Buffer
	//n, err := io.Copy(&buf, body)
	//if err != nil {
	//	fmt.Println("io.Copy err:", err)
	//	return
	//}
	//if n == 0 {
	//	fmt.Println("empty log")
	//	return
	//}
	//
	//err = body.Close()
	//if err != nil {
	//	fmt.Println("Close response body:", err)
	//}
}

//func recieveMessage(reader *bufio.Reader, message chan []byte) {
//	for {
//
//		message <- bytes
//	}
//}
