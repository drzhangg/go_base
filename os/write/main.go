package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

const (
	Format1 = "\b\b%d%%"
	Format2 = "\b\b\b%d%%"
	Format3 = "\b\b\b%d%%\b"
)

func progress(ch chan int64) {

	format := Format1

	var num int64 = 0

	for rate := range ch {
		if num > 10 && rate > 10 && rate < 100 {
			format = Format2
		} else if rate >= 100 {
			rate = 100

			format = Format3
		}
		fmt.Printf(format, rate)
		num = rate
	}
}

func main() {

	f, err := os.Open("/Users/drzhang/demo/go/go_base/os/write/file/22233509.jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rateCh := make(chan int64)
	defer close(rateCh)

	fmt.Print("rate:0%")

	go progress(rateCh)

	ret := make([]byte, 0)

	for {
		buf := make([]byte, 1024*5)
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		info, _ := f.Stat()

		if n == 0 {
			break
		}

		ret = append(ret, buf...)
		time.Sleep(time.Second)

		go func() {
			rateCh <- int64(len(ret)*100) / info.Size()
		}()

	}
	ioutil.WriteFile("/Users/drzhang/demo/go/go_base/os/write/rate.jpeg", ret, 0600)
}
