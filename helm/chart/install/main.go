package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/chart/loader"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
)

func main() {
	// 定义输入参数
	mysqlRootPassword := "default-password"
	mysqlUser := "default-user"
	mysqlPassword := "default-password"
	mysqlDatabase := "default-db"
	chartName := "azure-mirror/mysql"
	releaseName := "my-release"
	namespace := "default"

	// 生成 values 字段
	valuesMap := map[string]interface{}{
		"mysqlRootPassword": mysqlRootPassword,
		"mysqlUser":         mysqlUser,
		"mysqlPassword":     mysqlPassword,
		"mysqlDatabase":     mysqlDatabase,
	}

	// 将字段转换为 YAML 格式
	data, err := yaml.Marshal(&valuesMap)
	if err != nil {
		log.Fatalf("Error marshaling values to YAML: %v", err)
	}

	// 将 YAML 数据写入 values.yaml 文件
	filePath := "values.yaml"
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		log.Fatalf("Error writing values.yaml file: %v", err)
	}

	fmt.Printf("values.yaml file generated at %s\n", filePath)

	// 设置 Helm 配置
	settings := cli.New()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("Error initializing action configuration: %v", err)
	}

	// 创建 Helm install 操作
	install := action.NewInstall(actionConfig)
	install.ReleaseName = releaseName
	install.Namespace = namespace
	install.DryRun = true

	// 将 values.yaml 文件加载为 Helm values
	valueOpts := &values.Options{
		ValueFiles: []string{filePath},
	}

	p := getter.All(settings)

	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		log.Fatalf("Error merging values: %v", err)
	}

	// 安装 Helm chart
	chart, err := install.ChartPathOptions.LocateChart(chartName, settings)
	if err != nil {
		log.Fatalf("Error locating chart: %v", err)
	}

	load,err :=loader.Load(chart)
	if err != nil {
		log.Fatalf("Error load chart: %v", err)
	}

	rel, err := install.Run(load, vals)
	if err != nil {
		log.Fatalf("Error running Helm install: %v", err)
	}

	fmt.Printf("Successfully installed release: %s\n", rel.Name)
}

func getValuesYaml(valuesMap map[string]interface{}) ([]byte, error) {
	return yaml.Marshal(valuesMap)
}
