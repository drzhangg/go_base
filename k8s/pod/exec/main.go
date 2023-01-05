package main

import (
	"bytes"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/kubectl/pkg/scheme"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	h := os.Getenv("HOME")

	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("kubernetes.NewForConfig err:", err)
	}

	//clientset.CoreV1().

	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name("dev-14").
		Namespace("infra").
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Command: []string{"sh", "-c", "cmd"},
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		fmt.Println("remotecommand.NewSPDYExecutor err:", err)
		return
	}

	var stdout, stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  strings.NewReader(""),
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    true,
	})
	if err != nil {
		fmt.Println("exec.Stream err:", err)
	}

	//fmt.Println("stderr:", len(stdout.String()))
	//if err !=nil{
	//	fmt.Println("exec.Stream err:",err)
	//	return
	//}

	ret := map[string]string{"stdout": stdout.String(), "stderr": stderr.String(), "pod_name": "dev-14"}

	for {
		fmt.Println("ret::", ret)
		time.Sleep(time.Second)
	}

}
