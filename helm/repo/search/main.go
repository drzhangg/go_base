package main

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {
	settings := cli.New()

	// Initialize action configuration
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), "memory", log.Printf); err != nil {
		log.Fatalf("Error initializing action configuration: %v", err)
	}

	// Add the repository
	repoEntry := repo.Entry{
		Name: "azure-mirror",
		URL:  "http://mirror.azure.cn/kubernetes/charts/",
	}

	chartRepo, err := repo.NewChartRepository(&repoEntry, getter.All(settings))
	if err != nil {
		log.Fatalf("Error creating chart repository: %v", err)
	}

	_, err = chartRepo.DownloadIndexFile()
	if err != nil {
		log.Fatalf("Error downloading index file: %v", err)
	}

	// Initialize repository file
	repoFile := repo.NewFile()
	repoFile.Update(&repoEntry)

	// Write the repository file
	if err := repoFile.WriteFile(settings.RepositoryConfig, 0644); err != nil {
		log.Fatalf("Error writing repository file: %v", err)
	}

	// Load the repository index
	index, err := repo.LoadIndexFile(settings.RepositoryCache + "/azure-mirror-index.yaml")
	if err != nil {
		log.Fatalf("Error loading index file: %v", err)
	}

	type Description struct {
		Name string
		Describe string
		Version string
	}

	storeMap := make(map[string]Description)
	// Search for the chart in the index
	for name, charts := range index.Entries {
		//fmt.Printf("name:::%v\n",name)

		ch := charts[0]


		//fmt.Println(ch.Name,ch.Description,ch.Version,ch.AppVersion)

		storeMap[name] = Description{
			Name:     ch.Name,
			Describe: ch.Description,
			Version:  ch.AppVersion,
		}
	}
	fmt.Println("mm::",storeMap)
}
