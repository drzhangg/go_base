package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Jenkins 服务器的 URL
	jenkinsURL := "http://172.31.53.147:30180/"

	// Jenkins 用户名和 API Token（用于身份验证）
	username := "admin"
	password := "P@88w0rd"

	//jc := core.JenkinsCore{
	//	URL: jenkinsURL,
	//	UserName: username,
	//	Token: password,
	//}

	//jc.Request()


	// 构造请求 URL
	listURL := jenkinsURL + "/job/devops9b884/pipeline/api/json"

	// 创建 HTTP 请求
	request, err := http.NewRequest("GET", listURL, nil)
	if err != nil {
		panic(err)
	}

	// 进行基本身份验证
	request.SetBasicAuth(username,password)

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 打印响应内容
	fmt.Println(string(body))
}
