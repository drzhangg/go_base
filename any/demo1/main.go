package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"go_base/any/demo1/data"
)

const (
	DATABASE_Account = "DatabaseAccount"
	DATABASE_Schema = "DatabaseSchema"
)

var structInterfaceMap = map[string]data.TfeInterface{
	DATABASE_Account : &data.DatabaseAccount{},
	DATABASE_Schema : &data.DataBaseSchema{},
}

//var stuctMap = map[string]

func main() {
	da := data.DatabaseAccount{
		Basic: data.Basic{
			DeployStatus: "failed",
			RunId:        "123",
		},
		Name:  "jerry",
	}

	kind := "DatabaseAccount"
	
	req := data.Request{
		Status: "success1",
		Id:     "123abc",
	}

	datas,err := json.Marshal(&da)
	if err !=nil{
		fmt.Println("marshal err:",err)
	}

	in := data.NewInterface(kind,datas)

	in.UpdateStatus(req)

	fmt.Println(in)











}
