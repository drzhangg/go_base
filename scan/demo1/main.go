package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "testPrintColor", 0x1B)

	fmt.Printf("\n [1;32m%s%")


	fmt.Printf("\n %c[1;32m%s%c[m  %c[1;34m%s%c[m \n",0x1B,"hello",0x1B,0x1B,"logs",0x1B)


	a := "/app # u001b[6n"
	p:= "u001b[6n"

	//fmt.Println(strings.HasPrefix(a,"u001b[6n"))

	b := strings.ReplaceAll(a,p,"")

	r := strings.HasPrefix(a,p)
	fmt.Println(r)

	fmt.Println("b::",b)



	/*
	["{"Op":"stdout","Data":"lsrn","SessionID":"","Rows":0,"Cols":0}"]
	["{"Op":"stdout","Data":"u001b[1;32mhellou001b[m  u001b[1;34mlogsu001b[mrn/app # u001b[6n","SessionID":"","Rows":0,"Cols":0}"]

	*/
}