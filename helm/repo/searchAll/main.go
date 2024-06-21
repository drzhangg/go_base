package main

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {
	// 创建 Helm 的 CLI 环境
	settings := cli.New()

	// 获取 Helm 仓库配置文件路径
	repoFile := settings.RepositoryConfig

	// 加载 Helm 仓库配置文件
	repoFileContent, err := repo.LoadFile(repoFile)
	if err != nil {
		log.Fatalf("Failed to load repository file: %v", err)
	}

	// 检查是否有配置的仓库
	if len(repoFileContent.Repositories) == 0 {
		fmt.Println("No repositories found")
		return
	}

	// 更新所有仓库并打印所有 chart
	for _, re := range repoFileContent.Repositories {
		fmt.Printf("Updating repository: %s\n", re.Name)
		chartRepo, err := repo.NewChartRepository(re, getter.All(settings))
		if err != nil {
			log.Fatalf("Failed to create chart repository: %v", err)
		}
		if _, err := chartRepo.DownloadIndexFile(); err != nil {
			log.Fatalf("Failed to update repository index: %v", err)
		}

		// 加载仓库索引文件
		indexFilePath := settings.RepositoryCache + "/" + re.Name + "-index.yaml"
		indexFile, err := repo.LoadIndexFile(indexFilePath)
		if err != nil {
			log.Fatalf("Failed to load index file: %v", err)
		}

		// 打印所有 chart
		fmt.Printf("Charts in repository %s:\n", re.Name)
		for name, entries := range indexFile.Entries {
			for _, entry := range entries {
				fmt.Printf("  Name: %s, Version: %s, Description: %s\n", name, entry.Version, entry.Description)
			}
		}
	}
}
