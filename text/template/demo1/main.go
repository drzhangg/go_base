package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"text/template"
	"unicode"
)

type Params struct {
	Namespace  string      `json:"namespace"`
	Role       string      `json:"role"`
	EmptyFlag  bool        `json:"emptyFlag"`
	GroupRules []GroupRule `json:"groupRules"`
}

type GroupRule struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	KeyInformation string `json:"keyInformation"`
	Expr           string `json:"expr"`
	RuleType       string `json:"ruleType"`
	Email          string `json:"email"`
	Scope          string `json:"scope"`
	Zone           string `json:"zone"`
	ZoneName       string `json:"zoneName"`
	Id             string `json:"id"`
	MetricId       string `json:"metricId"`
	MetricName     string `json:"metricName"`
	Styles         string `json:"styles"`
	For            string `json:"for"`
	Threshold      string `json:"threshold"`
	Severity       string `json:"severity"`
	SelectAll      string `json:"selectAll"`
}

func main() {
	tmpl, err := template.ParseFiles("/Users/zhang/drzhang/demo/go/go_base/text/template/demo1/template/rule.yml")
	if err != nil {
		fmt.Println("err:", err)
	}

	param := &Params{}
	param.Role = "test-role"
	param.Namespace = "dev"

	param.GroupRules = []GroupRule{
		{
			Namespace:  "test",
			Name:       "test1",
			RuleType:   "test11",
			Email:      "test@qq.com",
			Scope:      "1",
			Zone:       "s-dev",
			ZoneName:   "s-dev",
			Id:         "123",
			MetricId:   "111",
			MetricName: "test-111",
			Styles:     "none",
			For:        "",
			Threshold:  "",
			Severity:   "",
			SelectAll:  "",
		},
		{
			Namespace:  "test",
			Name:       "test3",
			RuleType:   "test11",
			Email:      "test@qq.com",
			Scope:      "1",
			Zone:       "s-dev",
			ZoneName:   "s-dev",
			Id:         "123",
			MetricId:   "111",
			MetricName: "test-111",
			Styles:     "none",
			For:        "",
			Threshold:  "",
			Severity:   "",
			SelectAll:  "",
		},
		{
			Namespace:  "test",
			Name:       "test3",
			RuleType:   "test11",
			Email:      "test@qq.com",
			Scope:      "1",
			Zone:       "s-dev",
			ZoneName:   "s-dev",
			Id:         "123",
			MetricId:   "111",
			MetricName: "test-111",
			Styles:     "none",
			For:        "",
			Threshold:  "",
			Severity:   "",
			SelectAll:  "",
		},
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, param)
	if err != nil {
		fmt.Println("err1:", err)
	}

	//fmt.Println("buf:",string(buf.String()))

	var objList []*unstructured.Unstructured

	bufRd := bufio.NewReaderSize(&buf, 4096)
	//fmt.Println("bufRd:",bufRd)

	fmt.Println("isJsonArr:", isJsonArr(bufRd))

	if isJsonArr(bufRd) {

		err := json.NewDecoder(bufRd).Decode(&objList)
		if err != nil {
			fmt.Println("errr2:", err)
		}
	}else {
		decoder:=yaml.NewYAMLOrJSONDecoder(bufRd,4096)
		for{
			obj := &unstructured.Unstructured{}
			if err := decoder.Decode(obj);err != nil {
				if err == io.EOF{
					break
				}
				fmt.Println("decode err:",err)

			}
			objList = append(objList, obj)
		}
	}
	fmt.Println("objList:",&objList)

	for _,v := range objList{
		fmt.Println(v)
	}
}

var jsonArrPrefix = []byte("[")

func isJsonArr(rd *bufio.Reader) bool {
	size := 1024
	b, _ := rd.Peek(size)

	trim := bytes.TrimLeftFunc(b, unicode.IsSpace)
	return bytes.HasPrefix(trim, jsonArrPrefix)
}
