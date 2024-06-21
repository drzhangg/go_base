package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"os"

	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {
	// 初始化 Helm 设置
	settings := cli.New()

	// 定义 Chart 名称和版本
	chartName := "kubernetes-dashboard"
	chartVersion := "1.11.1"
	repoURL := "http://mirror.azure.cn/kubernetes/charts/"

	// 创建临时目录存放下载的 Chart
	tmpDir, err := os.MkdirTemp("", "helm-chart-")
	if err != nil {
		fmt.Printf("Failed to create temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	// 下载 Chart
	chartPath, err := downloadChart(repoURL, chartName, chartVersion, tmpDir, settings)
	if err != nil {
		fmt.Printf("Failed to download chart: %v\n", err)
		return
	}

	fmt.Printf("Chart downloaded to: %s\n", chartPath)
}

func downloadChart(repoURL, chartName, chartVersion, destDir string, settings *cli.EnvSettings) (string, error) {
	// 创建 Chart 存储库
	repoEntry := repo.Entry{
		Name: chartName,
		URL:  repoURL,
	}
	chartRepo, err := repo.NewChartRepository(&repoEntry, getter.All(settings))
	if err != nil {
		return "", fmt.Errorf("failed to create chart repository: %w", err)
	}

	// 下载 Chart 存储库的索引文件
	indexFilePath, err := chartRepo.DownloadIndexFile()
	if err != nil {
		return "", fmt.Errorf("failed to download index file: %w", err)
	}

	// 加载索引文件
	index, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to load index file: %w", err)
	}

	// 获取指定 Chart 版本信息
	chartVersionDetails, err := index.Get(chartName, chartVersion)
	if err != nil {
		return "", fmt.Errorf("failed to get chart version: %w", err)
	}

	// 使用 Helm 的下载器下载 Chart
	dl := downloader.ChartDownloader{
		Out:     os.Stdout,
		Getters: getter.All(settings),
		Options: []getter.Option{
			getter.WithURL(chartRepo.Config.URL),
		},
	}

	chartURL := chartVersionDetails.URLs[0]
	chartPath := fmt.Sprintf("%s-%s.tgz", chartName, chartVersion)
	destPath := fmt.Sprintf("%s/%s", destDir, chartPath)
	filename, _, err := dl.DownloadTo(chartURL, chartVersion, destDir)
	if err != nil {
		return "", fmt.Errorf("failed to download chart: %w", err)
	}

	fmt.Println("fileName:::",filename)

	chart, err := loader.Load(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to load chart: %v\n", err)
	}

	renderer := engine.Engine{}
	options := chartutil.ReleaseOptions{
		Name:      chartName,
		Namespace: "sre",
	}
	vals, err := chartutil.ToRenderValues(chart, nil, options, nil)
	if err != nil {
		return "", fmt.Errorf("Failed to create render values: %v\n", err)
	}

	renderedTemplates, err := renderer.Render(chart, vals)
	if err != nil {
		return "", fmt.Errorf("Failed to render template: %v\n", err)
	}

	fmt.Println("renderedTemplates:::",renderedTemplates)

	return destPath, nil
}
