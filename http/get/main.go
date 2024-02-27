package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//url := "http://127.0.0.1:8181/hello/shanghai"
	//req, err := http.Get(url)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//
	//datas, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//defer req.Body.Close()
	//
	//fmt.Println("strings:", string(datas))


	url := "http://127.0.0.1:8181/delete/zjh"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("newRequest err:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}


