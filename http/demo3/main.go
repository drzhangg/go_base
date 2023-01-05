package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	unix := time.Now().UnixMilli()

	url := fmt.Sprintf("t=%d", unix)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer resp.Body.Close()

	datas, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
	}

	err = ioutil.WriteFile("image.jpg",datas,0666)
	if err != nil {
		fmt.Println("WriteFile err:", err)
	}

}
