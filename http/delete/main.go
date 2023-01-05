package main

import (
	"fmt"
	"net/http"
)

func main() {
	delete()
}

func delete() {
	url := "http://127.0.0.1:8181/delete/zjh"
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("newRequest err:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	// handle error
	//}
	fmt.Println(resp.StatusCode)

	//fmt.Println(string(body))
}
