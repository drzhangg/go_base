package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/releaseutil"
	"os"
)

type DeployRequest struct {
	RepoURL      string                 // 仓库地址
	ChartName    string                 // Chart名称
	ChartVersion string                 // Chart版本
	Namespace    string                 // 命名空间
	ReleaseName  string                 // 在kubernetes中的程序名
	Values       map[string]interface{} // values.yaml 配置文件
}

func main() {

	env :=os.Getenv("HELM_DRIVER")
	fmt.Println("eee:",env)

	fmt.Println(installChart(&DeployRequest{
		RepoURL:      "http://mirror.azure.cn/kubernetes/charts/",
		ChartName:    "mysql",
		ChartVersion: "v1",
		Namespace:    "default",
		ReleaseName:  "test11",
		Values:       nil,
	}))
}

func installChart(deployRequest *DeployRequest) error {

	settings := cli.New()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(),deployRequest.Namespace,os.Getenv("HELM_DRIVER"),nil);err != nil {
		return fmt.Errorf("初始化 action 失败\n%s", err)
	}

	install := action.NewInstall(actionConfig)
	install.RepoURL = deployRequest.RepoURL
	install.Version = deployRequest.ChartVersion
	install.DryRun = true
	install.ClientOnly = true
	//install.Timeout = 30e9
	//install.CreateNamespace = true
	//install.Wait = true

	// k8s中的配置
	install.Namespace = deployRequest.Namespace
	install.ReleaseName = deployRequest.ReleaseName

	chartRequested,err := install.ChartPathOptions.LocateChart(deployRequest.ChartName,settings)
	if err != nil {
		return fmt.Errorf("下载失败\n%s", err)
	}

	chart,err := loader.Load(chartRequested)
	if err != nil {
		return fmt.Errorf("加载失败\n%s", err)
	}

	release1,err := install.Run(chart,nil)
	if err != nil {
		return fmt.Errorf("执行失败\n%s", err)
	}

	printReleaseYAML(release1)
	return nil
}

func printReleaseYAML(rel *release.Release)  {
	resources := releaseutil.SplitManifests(rel.Manifest)
	for _,resource := range resources{
		fmt.Println("---")
		fmt.Println(resource)
	}
}

