package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/repo"
	"log"
)

func main() {
	
	add(&repo.Entry{
		Name:                  "prometheus-community",
		URL:                   "https://prometheus-community.github.io/helm-charts",
		//Username:              "",
		//Password:              "",
		//CertFile:              "",
		//KeyFile:               "",
		//CAFile:                "",
		//InsecureSkipTLSverify: false,
		//PassCredentialsAll:    false,
	})

}

func add(entry *repo.Entry) error {
	settings := cli.New()

	repoFile := settings.RepositoryConfig

	repositories,err := repo.LoadFile(repoFile)
	if err != nil {
		repositories = repo.NewFile()
	}

	if repositories.Has(entry.Name){
		return  fmt.Errorf("仓库 %s 已存在",entry.Name)
	}

	repositories.Add(entry)

	if err = repositories.WriteFile(repoFile,0644); err != nil {
		return fmt.Errorf("无法保存仓库配置文件：%s", err)
	}
	

	log.Printf("成功添加仓库地址：%s。\n", entry.Name)
	return nil
}
