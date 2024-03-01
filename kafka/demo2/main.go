package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

type HttpRequest struct {
	// common request
	Action            string
	Zone              string
	Time_stamp        string
	Access_key_id     string
	Version           int
	Signature_method  string
	Signature_version int
	Signature         string
}

func main() {
	httpRequest := HttpRequest{
		Action:            "DescribeClusters",
		Zone:              "xining",
		Time_stamp:        time.Now().Format("2006-01-02T15:04:05Z"),
		//Time_stamp: fmt.Sprint("2024-02-27T11:09:10Z"),
		Access_key_id:     "GBPFITMISNISAOHQBYGC",
		Version:           1,
		Signature_method:  "HmacSHA256",
		Signature_version: 1,
	}

	data, err := structToSortedJSON(httpRequest)
	if err != nil {
		fmt.Println("err::", err)
	}

	fmt.Println("data::", string(data))

	params := bytesToHttpParams(data)
	fmt.Println("params:", params)

	secretAccessKey := "QxarSIh5sdB25RkjrvdHc0mcDS01Klrm3exJKD0I"

	signature := GetSignatureByParams(params,secretAccessKey)
	fmt.Println("signature:",signature)


	owner := "usr-zWBOsvlF"
	kafkaAppId := "app-n9ro0xcp"
	status := "active"


	encodeTs := url.QueryEscape(httpRequest.Time_stamp)
	//id := url.QueryEscape(fmt.Sprintf("apps.n[]=%s",kafkaAppId))
	//fmt.Println("iiiid:",id)

	pm1 := fmt.Sprintf("access_key_id=%s&action=DescribeClusters",httpRequest.Access_key_id)
	pm2 := fmt.Sprint("&apps.n%5B%5D=",kafkaAppId)
	//pm2 := fmt.Sprintf("&apps.n[]=%s",kafkaAppId)
	pm3 := fmt.Sprintf("&owner=%s&signature_method=%s&signature_version=%v&status.n=%s&time_stamp=%s&version=%d&zone=%s",
		owner,httpRequest.Signature_method,httpRequest.Signature_version,status,encodeTs,httpRequest.Version,httpRequest.Zone)


	pm := pm1+pm2+pm3
	fmt.Println("pm:",pm)

	si := GetSignatureByParams(pm,secretAccessKey)
	fmt.Println(si)

	newUrl := fmt.Sprint("https://xxx/iaas/?",pm,"&signature=",si)
	//url := "https://xxx/iaas/?access_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&apps.n%5B%5D=app-n9ro0xcp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&status.n=active&time_stamp=2024-02-27T11%3A09%3A10Z&version=1&zone=xining&signature=YHygQv7WC8gydT1dSnz58Cnwv9q9t1jadBU4urGZJ3g%3D"
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		fmt.Println("NewRequest err:",err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll err:",err)
	}

	//fmt.Println("bbbb:",string(body))

	dataMap := make(map[string]interface{})

	json.Unmarshal(body,&dataMap)

	fmt.Println("dm::",dataMap["action"])
	fmt.Println("dm::",dataMap["total_count"])

	qcks := []QingCloudKafkaStruct{}

	d1,err := json.Marshal(dataMap["cluster_set"])
	if err != nil {
		fmt.Println("json marshalll err:",err)
	}

	json.Unmarshal(d1,&qcks)

	//fmt.Println("ccc:",qcks)

	d11,_ :=json.Marshal(qcks)
	fmt.Println("d11:",string(d11))





	// "GET\n/iaas/\naccess_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&apps.n%5B%5D=app-n9ro0xcp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&status.n=active&time_stamp=2024-02-27T11%3A09%3A10Z&version=1&zone=xining"

	//YHygQv7WC8gydT1dSnz58Cnwv9q9t1jadBU4urGZJ3g%3D
	//YHygQv7WC8gydT1dSnz58Cnwv9q9t1jadBU4urGZJ3g%3D

}

func structToSortedJSON(s interface{}) ([]byte, error) {
	val := reflect.ValueOf(s)
	typ := val.Type()

	// 获取结构体字段名并排序
	var keys []string
	for i := 0; i < val.NumField(); i++ {
		keys = append(keys, typ.Field(i).Name)
	}
	sort.Strings(keys)

	// 创建一个新的 map 来存储排序后的字段和值
	sortedData := make(map[string]interface{})
	for _, key := range keys {
		field := val.FieldByName(key)

		if field.Interface() == "" {
			continue
		}
		sortedData[strings.ToLower(key)] = field.Interface()
	}

	// 转换为 JSON 格式
	return json.Marshal(sortedData)
}

func bytesToHttpParams(b []byte) string {
	var data map[string]interface{}

	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return ""
	}

	values := url.Values{}
	for key, value := range data {
		values.Add(key, fmt.Sprintf("%v", value))
	}

	httpParams := values.Encode()
	return httpParams
}

func GetSignatureByParams(params string, secretAccessKey string) string {

	//stringToSign := "GET\n/iaas/\naccess_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&apps.n%5B%5D=app-n9ro0xcp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&status.n=active&time_stamp=2024-02-27T11%3A09%3A10Z&version=1&zone=xining"
	//secretAccessKey := "QxarSIh5sdB25RkjrvdHc0mcDS01Klrm3exJKD0I"
	stringToSign := fmt.Sprintf("GET\n/iaas/\n%s", params)

	hn := hmac.New(sha256.New, []byte(secretAccessKey))
	hn.Write([]byte(stringToSign))
	sign := hn.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(sign)
	encodedSignature := url.QueryEscape(signature)

	fmt.Println("signature:", encodedSignature)
	return encodedSignature
}


// QingCloudKafkaStruct struct
type QingCloudKafkaStruct struct {
	AppId          string `json:"app_id"`
	CreateTime     string `json:"create_time"`
	Owner          string `json:"owner"`
	Name           string `json:"name"`
	NodeCount      int32  `json:"node_count"`
	ZoneId         string `json:"zone_id"`
	Status         string `json:"status"`
	Description    string `json:"description"`
	AppVersionInfo struct {
		Name        string `json:"name"`
		VersionType string `json:"version_type"`
		VersionId   string `json:"version_id"`
	} `json:"app_version_info"`
	ClusterId string `json:"cluster_id"`
}
