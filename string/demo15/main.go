package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "/checkJobName?value=devopsd6b9"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err:",err)
	}

	req.SetBasicAuth("admin","")

	res,err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	}

	datas, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
	}


	s := "<div/>"
	r := strings.EqualFold(s,string(datas))
	fmt.Println(r)

	//fmt.Println("strings:", string(datas))


}
