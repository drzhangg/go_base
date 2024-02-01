package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	var url string
	//var user,password string
	user := "root"
	password := "taosdata"

	host := "10.21.137.163:6041"
	db := "test"
	// http://%s:%s/rest/sql/%s
	url = fmt.Sprintf("http://%s/rest/sql/%s", host, db)

	auth := user+":"+  password
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	authHeader := "Basic " + authEncoded

	sqlstr := "select * from d1001"


	req, err := http.NewRequest("POST", url, strings.NewReader(sqlstr))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	fmt.Println("body::",string(body))

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	if result["code"].(float64) == 0 {
		dataValue := []interface{}{}
		for _, v := range result["data"].([]interface{}) {
			dataValue = append(dataValue, v.([]interface{})[0])
		}
		fmt.Println("data value:::",dataValue)
	} else {
		fmt.Println("desccccc:", result["desc"])
		//fmt.Println(result["desc"])
	}
}
