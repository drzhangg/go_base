package main

import (
	"fmt"
	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/contact/user"
	"net/url"
)

type WeCom struct {
	app *corporation.App
}

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Parentid int    `json:"parentid"`
	Order    int    `json:"order"`
	Users []User
}

func NewWeCom() *WeCom {
	Corp := corporation.New(corporation.Config{Corpid: "xx"})

	app := Corp.NewApp(corporation.AppConfig{
		AgentId:        "xx",
		Secret:         "xx-z-xx",
		Token:          "xx",
		EncodingAESKey: "xxx",
	})

	return &WeCom{
		app: app,
	}
}

type User struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	UserID string `json:"userid"`
	Email  string `json:"email"`
}

func main() {
	w := NewWeCom()

	params := url.Values{}
	params.Add("department_id","235")

	userList,err := user.List(w.app,params)
	if err != nil {
		fmt.Println("list user err:",err)
		return
	}

	fmt.Println("user list::",string(userList))

	//resp, err := department.List(w.app, params)
	//if err != nil {
	//	fmt.Println("list err:",err)
	//	return
	//}

	//resp1,err := user.List(w.app,params)
	//if err != nil {
	//	fmt.Println("users err:",err)
	//	return
	//}
	//fmt.Println("users:",string(resp1))
	//
	//dept := struct {
	//	Department []Department `json:"department"`
	//}{}
	//
	//err = json.Unmarshal(resp, &dept)
	//if err != nil {
	//	fmt.Println("unmarshal err:",err)
	//}
	//
	//fmt.Println("dept::",string(resp))

	//for _, val := range dept.Department{
	//	params.Add("department_id", strconv.Itoa(val.ID))
	//	fmt.Println("val::",val)
	//}

}
