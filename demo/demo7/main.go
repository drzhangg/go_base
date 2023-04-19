package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

type Data struct {
	typeModel   string
	author      string
	project     string
	projectUrl  string
	projectId   string
	releateNote string
}

type typeMode map[string][]Data

func main() {
	fmt.Println(parseUrl("https://github.com/test"))

	file := "xxx+Enterprise+v3.4+Changelog.xlsx"
	path := "/Users/drzhang/Downloads/"
	f, err := excelize.OpenFile(path + file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	typeMap := typeMode{}
	modelMap := make(map[string]typeMode)

	//typeMap

	//modelMap := make(map[string]Data)

	for index, _ := range rows {

		//if row[8]!= ""{
		//	modelMap[row[8]] = append(modelMap[row[8]], row[7])
		//}

		if len(rows[index]) >= 9 {
			data := Data{}

			projectUrl := rows[index][0]
			data.projectUrl = projectUrl
			//fmt.Println(rows[index][0])

			project, projectId := parseUrl(projectUrl)
			if project != "" && projectId != "" {
				data.project = project
				data.projectId = projectId
			}

			data.author = rows[index][2]
			data.releateNote = rows[index][6]

			//fmt.Println("roww:",rows[index][7])

			typeMap[rows[index][7]] = append(typeMap[rows[index][7]], data)
			modelMap[rows[index][8]] = typeMap
			//fmt.Println("row:",row[8])

		}
	}
	//fmt.Println(modelMap)
	//fmt.Println(typeMap)

	for k, v := range modelMap {
		if k == "DevOps" {
			for k1, v1 := range v {
				fmt.Println("## ", k)
				fmt.Println("### ", k1)

				for i := 0; i < len(v1); i++ {
					fmt.Printf("- %s [%s#%s](%s) by [@%s]\n", v1[i].releateNote, v1[i].project, v1[i].projectId, v1[i].projectUrl, v1[i].author)

				}
			}
		}
	}
}

func parseUrl(url string) (string, string) {
	urls := strings.Split(url, "//")

	if len(urls) >= 2 {
		arr := strings.Split(urls[1], "/")
		if len(arr) == 5 {
			return arr[2], arr[4]
		}
	}
	return "", ""
}

/*
## module1

### type1

- release note 1 [project#pr_id](pr_url) by [@author](pr_url)
- release note 2 [project#pr_id](pr_url) by [@author](pr_url)

- %s [%s#%s](%s) by [@%s]

- release note 1 [kse-console#32](https://github.com/xxx/xxx/pull/32) by [@harrisonliu5](https://github.com)

## module1

### type2

- release note 1 [project#pr_id](pr_url) by [@author](pr_url)
- release note 2 [project#pr_id](pr_url) by [@author](pr_url)

## module2

### type1

- release note 1 [project#pr_id](pr_url) by [@author](pr_url)
- release note 2 [project#pr_id](pr_url) by [@author](pr_url)

*/
