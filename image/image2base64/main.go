package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {

	ff, err := ioutil.ReadFile("/Users/zhang/drzhang/demo/go/go_base/image.jpg")
	if err !=nil{
		fmt.Println("err:",err)
	}

	baseImg := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(ff)               // 文件转base64

	fmt.Println(baseImg)

}
