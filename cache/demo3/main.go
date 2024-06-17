package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(0,0)

	c.Set("name","jerry",cache.NoExpiration)


	name,found := c.Get("name")
	if found{
		fmt.Println("name::",name)
	}
}
