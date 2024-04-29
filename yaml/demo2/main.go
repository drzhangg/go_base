package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gopkg.in/yaml.v3"
)

type authorities struct {
	Authorities []Data
}
type Data struct {
	Role string
	Verbs map[string]map[string]bool
}

func main() {
	c := config.New(
		config.WithSource(
			file.NewSource("/Users/drzhang/demo/go/go_base/yaml/demo2/role.yaml"),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			configData := kv.Value
			//configData = uConf.ReplaceEnv(configData)
			return yaml.Unmarshal(configData, v)
		}),
	)
	if err := c.Load(); err != nil {
		fmt.Println("c load failed:",err)
		return
	}

	var roles authorities
	if err := c.Scan(&roles); err != nil {
		panic(err)
	}

	fmt.Println("role::",roles)

}
