package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/repo"
	"log"
)

// DeployRequest
/**
 * 部署时用到的结构体
 */
type DeployRequest struct {
	RepoURL      string                 // 仓库地址
	ChartName    string                 // Chart名称
	ChartVersion string                 // Chart版本
	Namespace    string                 // 命名空间
	ReleaseName  string                 // 在kubernetes中的程序名
	Values       map[string]interface{} // values.yaml 配置文件
}

// ---------------------------------------------------------------

// ChartListResponse
/**
 * 返回指定仓库中的所有Chart信息
 */
type ChartListResponse struct {
	ChartName    string // Chart名称
	ChartVersion string // Chart版本
	AppVersion   string // 应用版本
	Description  string // 描述
}

func main() {
	//settings := cli.New()
	//actionConfig := new(action.Configuration)
	//client := action.NewInstall(actionConfig)
	//
	//chartName := "stable/mysql"
	//releaseName := "my-mysql"
	//
	//client.Namespace = "default"
	//client.ReleaseName = releaseName
	//
	//rel, err := client.Run(chartName, nil)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error installing chart: %s\n", err.Error())
	//	os.Exit(1)
	//}
	//
	//fmt.Printf("Successfully installed chart. Release name: %s\n", rel.Name)

	//repoEntry,err := list()
	//if err != nil {
	//	fmt.Println("err:",err)
	//}
	//
	//for _, item := range repoEntry{
	//	fmt.Println("name::",item.Name)
	//	fmt.Println("Password::",item.Password)
	//	fmt.Println("URL::",item.URL)
	//	fmt.Println("string::",item.String())
	//}

	res,err := searchAll("kong","")
	if err != nil {
		fmt.Println("err:",err)
	}

	for _,v := range res{
		fmt.Println(v.ChartName)
		fmt.Println(v.Description)
		fmt.Println(v.ChartVersion)
		fmt.Println(v.AppVersion)
	}
}

// 查看仓库信息
func list() ([]*repo.Entry, error) {
	settings := cli.New()

	// 加载仓库配置文件
	repositories, err := repo.LoadFile(settings.RepositoryConfig)
	if err != nil {
		return nil, fmt.Errorf("无法保存仓库配置文件：%s", err)
	}
	return repositories.Repositories, nil
}


// searchAll(仓库名, Chart名)
func searchAll(repoName, chartName string) ([]*ChartListResponse, error) {
	settings := cli.New()

	path := fmt.Sprintf("%s/%s-index.yaml", settings.RepositoryCache, repoName)
	// 加载 xxx-index.yaml 文件
	indexFile, err := repo.LoadIndexFile(path)
	if err != nil {
		return nil, fmt.Errorf("仓库 %s 不存在", repoName)
	}

	var chartList []*ChartListResponse

	// 遍历指定仓库的 Chart 信息
	for _, entry := range indexFile.Entries[chartName] {
		// 将每个 Chart 的主要信息提取出来
		chart := &ChartListResponse{
			ChartName:    entry.Name,
			ChartVersion: entry.Version,
			AppVersion:   entry.AppVersion,
			Description:  entry.Description,
		}
		chartList = append(chartList, chart)
	}

	// 指定仓库的Chart信息

	//logger.Debugf("%s", chartList)

	log.Printf("%s",chartList)
	return chartList, nil
}