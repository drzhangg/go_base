package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ff, err := ioutil.ReadFile("/Users/zhang/drzhang/demo/go/go_base/image.jpg")
	if err !=nil{
		fmt.Println("err:",err)
	}

	baseImg := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(ff)

	url := "https://web.baimiaoapp.com/api/ocr/image/xunfei"

	m := make(map[string]interface{})
	m["dataUrl"] = baseImg
	m["name"] = "image.jpg"
	m["isSuccess"] = "false"
	m["batchId"] = ""
	m["hash"] = "false"
	m["isSuccess"] = "false"
	m["isSuccess"] = "false"
	m["isSuccess"] = "false"

	data,err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal err:",err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
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
