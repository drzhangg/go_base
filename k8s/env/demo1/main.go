package main

import (
	"fmt"
	"k8s.io/client-go/rest"
)

func main() {
	cfg,err := rest.InClusterConfig()
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
