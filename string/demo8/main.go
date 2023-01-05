package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Phone string
	A int
}

type Name struct {
	Name string
}

type Age struct {
	Age int
}

type Address struct {
	Address string
}

func main() {
	//u := User{"12312312qweqw",22}
	name := Name{"jerry"}
	age := Age{25}
	add := Address{"上海"}

	//ud, err := strToJson(u)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}

	nd, err := strToJson(name)
	if err != nil {
		fmt.Println("err:", err)
	}
	ad, err := strToJson(age)
	if err != nil {
		fmt.Println("err:", err)
	}
	addd, err := strToJson(add)
	if err != nil {
		fmt.Println("err:", err)
	}

	s := []string{string(nd), string(ad), string(addd)}
	//s := []string{string(ud),string(nd), string(ad), string(addd)}
	fmt.Println(string(send(s...)))

	fmt.Println(strings.Trim(string(send(s...)),"a[]"))
	ts := strings.Trim(string(send(s...)),"a[]")

	splits := strings.Split(ts,",")
	fmt.Println("sss:",splits)

	var rs string
	err = json.Unmarshal([]byte(splits[0]),&rs)
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println("rs:",rs)

	resultname := Name{}
	json.Unmarshal([]byte(rs),&resultname)
	fmt.Println("resultName:",resultname)


}

func strUnmarshal(b []byte)  {
	js := string(b[3:len(b)-2])
	fmt.Println(js)


}

func quote(in string) string {

	quoted, _ := json.Marshal(in)
	return string(quoted)
}

func transform(values []string, transformFn func(string) string) []string {
	ret := make([]string, len(values))
	for i, msg := range values {
		ret[i] = transformFn(msg)
	}
	return ret
}

func send(messages ...string) []byte {
	if len(messages) > 0 {
		return []byte(fmt.Sprintf("a[%s]", strings.Join(transform(messages, quote), ",")))
	}
	return nil
}

func strToJson(s interface{}) ([]byte, error) {
	data, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}
	return data, nil
}
