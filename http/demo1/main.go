package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var (
	t = cache.New(time.Minute*10, time.Minute*15)
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {


		name, ok := t.Get("name")
		if ok {
			fmt.Println("get ok")
			w.Write([]byte("name::"+name.(string)))
			return
		}

		t.Set("name", "jerry",cache.NoExpiration)
		w.Write([]byte("get name failed"))
	})

	http.ListenAndServe(":9999", nil)
}
