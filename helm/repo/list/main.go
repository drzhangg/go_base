package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {

	fmt.Println(list())

}

func list() ([]*repo.Entry, error) {
	settings := cli.New()

	repositories,err := repo.LoadFile(settings.RepositoryConfig)
	if err != nil {
		return nil, fmt.Errorf("无法保存仓库配置文件：%s", err)
	}
	return repositories.Repositories,nil
}
