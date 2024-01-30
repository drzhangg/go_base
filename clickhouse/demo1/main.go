package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mailru/go-clickhouse/v2"
	"log"
)

func main() {
	dataSouce := fmt.Sprintf("http://%s", "127.0.0.0")
	fmt.Println(dataSouce)

	connect, err := sql.Open("chhttp", "http://xxx:8123/test")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal("ping err ",err)
	}

	sqlstr := "select name from test1 where name = ? "

	//var name string
	//err = connect.QueryRow(sqlstr, 1).Scan(&name)
	//if err != nil {
	//	log.Fatal("QueryRow err ",err)
	//}
	//
	//fmt.Println("name:",name)


	rows,err := connect.Query(sqlstr,"jerry")
	if err != nil {
		log.Fatal("QueryRow err ",err)
	}

	if !rows.Next(){
		fmt.Println("no data")
	}

	var value string
	if err := rows.Scan(&value); err != nil {
		log.Println("scan err:",err)
		//return false, err.Error()
	}
	fmt.Println("val::",value)

}
