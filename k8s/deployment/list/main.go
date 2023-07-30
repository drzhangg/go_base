package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/transport"
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

	conf := &transport.Config{
		UserAgent:          config.UserAgent,
		Transport:          config.Transport,
		WrapTransport:      config.WrapTransport,
		DisableCompression: config.DisableCompression,
		TLS: transport.TLSConfig{
			Insecure:   config.Insecure,
			ServerName: config.ServerName,
			CAFile:     config.CAFile,
			CAData:     config.CAData,
			CertFile:   config.CertFile,
			CertData:   config.CertData,
			KeyFile:    config.KeyFile,
			KeyData:    config.KeyData,
			NextProtos: config.NextProtos,
		},
		Username:        config.Username,
		Password:        config.Password,
		BearerToken:     config.BearerToken,
		BearerTokenFile: config.BearerTokenFile,
		Impersonate: transport.ImpersonationConfig{
			UserName: config.Impersonate.UserName,
			Groups:   config.Impersonate.Groups,
			Extra:    config.Impersonate.Extra,
		},
		Dial:  config.Dial,
		Proxy: config.Proxy,
	}

	fmt.Printf("%#v\n",conf)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	deploys,err := clientset.AppsV1().Deployments("").List(context.TODO(),v1.ListOptions{
		//LabelSelector: "control-plane",
	})
	if err != nil {
		fmt.Println("get deploy list err:", err)
	}

	for _,v := range deploys.Items{
		fmt.Println(v.Name)
	}

}
