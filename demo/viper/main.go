package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.AddConfigPath("/Users/zhang/drzhang/demo/go/go_base/demo/viper")
	v.SetConfigName("configa")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	name := v.GetString("data.name")
	//s := v.GetStringSlice("name.name")
	fmt.Println(name)
}
