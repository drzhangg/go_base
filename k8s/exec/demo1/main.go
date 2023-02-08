package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/kubectl/pkg/scheme"
	"os"
	"path/filepath"
)

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config.conf")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	req := clientset.CoreV1().RESTClient().Post().Resource("pods").Name("nginx-test-6c88c44876-vf9vs").Namespace("default").SubResource("exec")
	req.VersionedParams(&v1.PodExecOptions{
		Container: "nginx",
		Command:   []string{"bash"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err.Error())
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    false,
		//TerminalSizeQueue: nil,
	})
	if err != nil {
		panic(err.Error())
	}

	// 返回数据
	//ret := map[string]string{"stdout":os, "stderr":stderr.String(), "pod_name": podName}
	//return ret, nil
}
