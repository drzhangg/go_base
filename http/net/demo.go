package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {


	addr := "https://onlineservices.immigration.govt.nz/Registration/default.aspx"


	user := struct {
		Gender  int    `json:"gender"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}{
		Gender: 1,
		Age: 26,
		Address: "shanghai",
	}

	data, err := json.Marshal(&user)
	if err != nil {
		fmt.Println("marshal err:", err)
	}

	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(data))
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

	fmt.Println(string(body))
}
