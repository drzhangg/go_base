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

func main() {

	//{"status":[],"verbose":1,"clusters":["cl-sz2chfvp"],"action":"DescribeClusters","owner":"usr-zWBOsvlF","zone":"xn1a"}
	hr := HttpRequest{
		Action:            "DescribeClusters",
		Zone:              "xn1a",
		Time_stamp:        fmt.Sprintf(time.Now().Format("2006-01-02T15:04:05Z")),
		Access_key_id:     "GBPFITMISNISAOHQBYGC",
		Version:           1,
		Signature_method:  "HmacSHA256",
		Signature_version: 1,
		Clusters:           []string{"cl-sz2chfvp"},
		Owner:             "usr-zWBOsvlF",
		Verbose: 1,
	}

	pm1 := fmt.Sprintf("access_key_id=%s&action=DescribeClusters",hr.Access_key_id)
	pm2 := fmt.Sprint("&clusters.n%5B%5D=",hr.Clusters[0])
	//pm2 := fmt.Sprintf("&apps.n[]=%s",kafkaAppId)
	pm3 := fmt.Sprintf("&owner=%s&signature_method=%s&signature_version=%d&time_stamp=%s&verbose=%d&version=%d&zone=%s",
		hr.Owner,hr.Signature_method,hr.Signature_version, url.QueryEscape(hr.Time_stamp),hr.Verbose,hr.Version,hr.Zone)


	pm := pm1+pm2+pm3

	secretAccessKey := "QxarSIh5sdB25RkjrvdHc0mcDS01Klrm3exJKD0I"
	signature := getSignatureByParams(pm,secretAccessKey)

	//fmt.Println("signature:",signature)

	newUrl := fmt.Sprint("https://xxx/iaas/?", pm, "&signature=", signature)
	fmt.Println("newUrl::",newUrl)
	data, _, err := Get(newUrl)
	if err != nil {
		fmt.Printf("get qingcloud kafka resource failed, err: %v", err)
	}

	//fmt.Println("data::",string(data))

	dataMap := make(map[string]interface{})
	qckns := QingCloudKafkaStruct{}

	json.Unmarshal(data,&dataMap)

	dd := dataMap["cluster_set"].([]interface{})[0].(map[string]interface{})

	nodeData,err := json.Marshal(dd)
	if err != nil {
		fmt.Println("marsha err:",err)
	}

	json.Unmarshal(nodeData,&qckns)

	d111,_ :=json.Marshal(qckns)
	fmt.Println("11:",string(d111))

	//fmt.Printf("qcccc:%#v\n",qckns)
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
	ClusterId string                     `json:"cluster_id"`
	Nodes     []QingCloudKafkaNodeStruct `json:"nodes"`
}

type QingCloudKafkaNodeStruct struct {
	ClusterId string `json:"cluster_id"`
	NodeId    string `json:"node_id"`
	Role      string `json:"role"`
	PrivateIp string `json:"private_ip"`
	Status    string `json:"status"`
}

type HttpRequest struct {
	// common request
	Action            string
	Zone              string
	Time_stamp        string
	Access_key_id     string
	Version           int
	Signature_method  string
	Signature_version int
	//Signature         string
	Clusters []string
	Owner   string
	Verbose int
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


func Get(url string) ([]byte, int, error) {
	//url := "https://xxx/iaas/?access_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&apps.n%5B%5D=app-n9ro0xcp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&status.n=active&time_stamp=2024-02-27T11%3A09%3A10Z&version=1&zone=xining&signature=YHygQv7WC8gydT1dSnz58Cnwv9q9t1jadBU4urGZJ3g%3D"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("the url:%s, get failed, err:%v\n", url, err)
		return nil, 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("the url:%s, read body failed, err:%s", url, err)
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func getSignatureByParams(params string, secretAccessKey string) string {

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
