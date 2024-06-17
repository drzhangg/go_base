package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {
	update()
}

func update() (string, error) {
	settings := cli.New()

	repositories,err := repo.LoadFile(settings.RepositoryConfig)
	if err != nil {
		return "", fmt.Errorf("无法加载仓库配置文件：%s", err)
	}

	for _, repoEntry := range repositories.Repositories{
		chartRepository,err := repo.NewChartRepository(repoEntry,getter.All(settings))
		if err != nil {
			return "", fmt.Errorf("无法添加仓库：%s\n", err)
		}

		// 更新仓库索引信息
		if _, err := chartRepository.DownloadIndexFile(); err != nil {
			return "", fmt.Errorf("无法下载仓库索引：%s\n", err)
		}
		fmt.Printf("...Successfully got an update from the %s chart repository\n", repoEntry.Name)

	}
	return "Update Complete. ⎈Happy Helming!⎈", nil
}