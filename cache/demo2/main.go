package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {

	t := cache.New(time.Second*10,time.Second*15)

	t.SetDefault("age","333")

	for {
		age,ok := t.Get("age")
		if !ok {
			break
		}

		fmt.Println("name::", age, ", time:",time.Now())
		time.Sleep(time.Second)
	}


}
