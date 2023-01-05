package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://web.baimiaoapp.com/api/perm/single"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	req.Header.Add("x-auth-token","hUloQXv2CION4a9pmxftAS3Hy96nHcKDy6qsqA3jC9foQZNvWXsEDueXmzglPPPR")
	req.Header.Add("x-auth-uuid","a621f5e7-d378-46b0-8e68-3feb2912a580")

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
	fmt.Println("data:",string(datas))

}
