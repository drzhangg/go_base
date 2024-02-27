package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
)

type User struct {
	Name string
	Age int
	Address string
	Color string
}

type JsonObj map[string]interface{}

func (j JsonObj) MarshalJSON() ([]byte, error) {
	// 获取map的keys并排序
	keys := make([]string, 0, len(j))
	for k := range j {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 创建一个 buffer 来存储字节
	var buf bytes.Buffer

	buf.WriteByte('{')
	for i, k := range keys {
		if i != 0 {
			buf.WriteByte(',')
		}
		// 写入key
		key, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		buf.Write(key)

		buf.WriteByte(':')

		// 写入value
		val, err := json.Marshal(j[k])
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}
	buf.WriteByte('}')

	return buf.Bytes(), nil
}

func main() {
	u := User{
		Name:    "drzhangg",
		Age:     20,
		Address: "wuhan",
		Color:   "blue",
	}

	bb,_ := json.Marshal(&u)

	j1 := JsonObj{}

	err := json.Unmarshal(bb,&j1)
	if err != nil {
		fmt.Println("error1:", err)
	}

	fmt.Println("j1::",j1)


	j := JsonObj{"name": "John", "age": 30, "city": "New York","acc":"sda"}
	b, err := json.Marshal(j)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))


	var data map[string]interface{}
	err = json.Unmarshal([]byte(b), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	values := url.Values{}
	for key, value := range data {
		values.Add(key, fmt.Sprintf("%v", value))
	}

	httpParams := values.Encode()
	fmt.Println(httpParams)

	//str1 := strings.Replace(string(b),",","&",-1)
	//
	//str2 := strings.Replace(str1,":","=",-1)
	//
	//str3 := strings.Replace(str2,`"`,"",-1)
	//str4 := strings.Replace(str2,`"`,"",-1)
	//
	//fmt.Println("str1",str3)
	//for key,v := range j{
	//	fmt.Println(key,"::;",v)
	//}
}

