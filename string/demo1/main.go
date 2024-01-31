package main

import "fmt"

func main() {
	// u001b[1;32mhellou001b[m  u001b[1;34mlogsu001b[m
	//a := "u001b[1;32mhellou001b[m  u001b[1;34mlogsu001b[m"
	//
	//b := strings.ReplaceAll(a,"u001b","%c")
	//
	//ss := []interface{}{}
	//for i := 0; i < strings.Count(a,"u001b"); i++ {
	//	ss = append(ss, 0x1B)
	//}
	//fmt.Printf(b,ss...)

	flag := false

	dbName := "test"
	dsn := fmt.Sprintf("clickhouse://%s", "127.0.0.1:8123")

	if dbName != "" {
		dsn = fmt.Sprintf(dsn+"?database=%s", dbName)
	}

	if flag {
		dsn = fmt.Sprintf(dsn+"&username=%s&password=%s", "admin","admin1")
	}




	fmt.Println(dsn)
}
