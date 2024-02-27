package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Person struct {
	Name string
	Age  int
	City string
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
		sortedData[strings.ToLower(key)] = field.Interface()
	}

	// 转换为 JSON 格式
	return json.Marshal(sortedData)
}

func main() {
	p := Person{Name: "Alice", Age: 30, City: "New York"}
	sortedJSON, err := structToSortedJSON(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(sortedJSON))
}
