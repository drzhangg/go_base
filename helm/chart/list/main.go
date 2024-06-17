package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"os"
)

func main() {
	namespace := "default" // 设置命名空间

	// 加载 kubeconfig 配置
	//kubeconfig := os.Getenv("KUBECONFIG")
	settings := cli.New()

	//config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	//if err != nil {
	//	fmt.Printf("Error building kubeconfig: %v\n", err)
	//	os.Exit(1)
	//}

	// 创建 Helm 的 action 配置
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), debug); err != nil {
		fmt.Printf("Error initializing action configuration: %v\n", err)
		os.Exit(1)
	}

	// 列出所有 releases
	client := action.NewList(actionConfig)
	releases, err := client.Run()
	if err != nil {
		fmt.Printf("Error listing releases: %v\n", err)
		os.Exit(1)
	}

	for _, rel := range releases {
		fmt.Printf("NAME: %s, NAMESPACE: %s, REVISION: %d, UPDATED: %s, STATUS: %s, CHART: %s, APP VERSION: %s\n",
			rel.Name, rel.Namespace, rel.Version, rel.Info.LastDeployed, rel.Info.Status, rel.Chart.Name(), rel.Chart.AppVersion())
	}
}

func debug(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
