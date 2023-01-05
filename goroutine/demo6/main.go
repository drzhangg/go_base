package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {

	list := []string{"123", "321", "abc", "sdada", "qweqeqw"}
	for i := 0; i < 100000; i++ {
		list = append(list, fmt.Sprintf("hello %v", i))
	}

	ch := make(chan struct{}, 10)
	defer close(ch)

	eg, _ := errgroup.WithContext(context.TODO())

	for _, v := range list {
		ch <- struct{}{}
		val := v
		eg.Go(func() error {
			defer func() {
				<-ch
			}()

			if val == "abc" {
				//return errors.New("have error")
			} else {
				fmt.Println(val)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
}
