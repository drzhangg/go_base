package main

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"log"
)

func main() {
	dsn := "clickhouse://xxx:8123?username=default&password=default&database=test"
	// 打开数据库连接
	conn, err := sql.Open("clickhouse", dsn)
	if err != nil {
		fmt.Println("连接失败:", err)
		log.Fatal("Open err ",err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		fmt.Println("ping失败:", err)
		log.Fatal("Ping err ",err)
	}


	sqlstr := "select name from test1 where name = ? "

	rows,err := conn.Query(sqlstr,"jerry")
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
