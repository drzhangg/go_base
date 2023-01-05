package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type test struct {
	scope *cache.Cache
}

func main() {
	t := new(test)
	t.scope = cache.New(time.Second*10, time.Second*20)

	t.scope.Set("name", "jerry", time.Second*20)


	for {
		name, ok := t.scope.Get("name")
		if !ok {
			break
		}

		fmt.Println("name::", name, ", time:",time.Now())
		time.Sleep(time.Second)
	}

}
