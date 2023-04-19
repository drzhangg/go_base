package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type Data struct {
	typeModel  string
	author     string
	project    string
	projectUrl string
	projectId  int
}

type typeMode map[string][]string

func main() {
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

	for index, _ := range rows {

		//if row[8]!= ""{
		//	modelMap[row[8]] = append(modelMap[row[8]], row[7])
		//}

		if len(rows[index]) >= 9 {
			//fmt.Println("roww:",rows[index][7])

			typeMap[rows[index][7]] = append(typeMap[rows[index][7]], rows[index][6])
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
					if v1[i] != "" {
						fmt.Println("- ", v1[i])
					}

				}
			}
		}
	}
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
