package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://127.0.0.1:8181/hello/shanghai"
	req, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	datas, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer req.Body.Close()

	fmt.Println("strings:", string(datas))
}


