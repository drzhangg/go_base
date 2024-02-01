package inspector

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TdengineConfig struct {
	Host     string            `json:"host"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Data     map[string]string `json:"data"`
}

func (i *TdengineConfig) Inspect(ctx context.Context) (ok bool, res string) {
	// todo
	//if len(i.Config.Data) > 0 {
	//	return i.ProbeWithData()
	//}

	return i.ProbeWithPing()
}

func (i *TdengineConfig) ProbeWithPing() (ok bool, res string) {
	return i.RunSql()
}

func (i *TdengineConfig) RunSql() (bool, string) {
	var db string
	url := fmt.Sprintf("http://%s/rest/sql/%s", i.Host, db)

	auth := i.Username + ":" + i.Password
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	authHeader := "Basic " + authEncoded

	// sql
	sqlstr := "select server_version()"

	req, err := http.NewRequest("POST", url, strings.NewReader(sqlstr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, fmt.Sprintf("Create tdengin data request failed :%s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, fmt.Sprintf("Request tdengin data failed :%s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return false, fmt.Sprintf("Reading tdengin data failed :%s", err.Error())
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return false, fmt.Sprintf("Get tdengin data failed :%s", err.Error())
	}

	if result["code"].(float64) != 0 {
		return false, "Ping tdengine failed!"
	}

	return true, "Ping tdengine success!"
}
