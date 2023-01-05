package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Quantity struct {
	DevProd   int `json:"dev-prod"`
	Dev       int `json:"dev"`
	B2cDev    int `json:"b2c-dev"`
	Container int `json:"container"`
}

func main() {
	pod()
	deploy()
}

func pod() {
	//ds := []string{"dev","dev-prod","b2c-dev"}
	//for _,v := range ds{
	//
	//}

	namespaces := []string{"123-dev-prod", "asda-dev", "111-b2c-dev", "231-dev", "1312312312", "a-dev-prod", "00123-dev"}

	var zoneMap = returnMap()
	nm := make(map[string]int)
	for _, v := range namespaces {
		nm = splitZone(zoneMap, v)
	}
	fmt.Println(nm)

	data, err := json.Marshal(&nm)
	if err != nil {
		fmt.Println("err:", err)
	}
	q := Quantity{}
	err = json.Unmarshal(data, &q)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("q:", q)
}

func deploy() {
	namespaces := []string{"123-dev", "asda-dev", "111", "231-dev-prod", "1312312312", "a-dev-prod", "00123-dev"}

	var zoneMap = returnMap()
	nm := make(map[string]int)
	for _, v := range namespaces {
		nm = splitZone(zoneMap, v)
	}
	fmt.Println(nm)

	data, err := json.Marshal(&nm)
	if err != nil {
		fmt.Println("err:", err)
	}
	q := Quantity{}
	err = json.Unmarshal(data, &q)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("q:", q)
}

func mapMarshal(m map[string]int, d interface{}) {

}

func returnMap() map[string]int {
	return map[string]int{
		"dev-prod":  0,
		"dev":       0,
		"b2c-dev":   0,
		"container": 0,
	}
}

func splitZone(m map[string]int, namespace string) map[string]int {
	splitNs := strings.Split(namespace, "-")
	if len(splitNs) > 1 {
		splitNs = splitNs[1:]
		m[strings.Join(splitNs, "-")]++
	} else {
		m["container"]++
	}
	return m
}
