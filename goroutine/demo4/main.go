package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	list := []string{"123","321","abc","sdada","qweqeqw"}

	eg,ctx := errgroup.WithContext(context.Background())

	ch := make(chan struct{},4)
	defer close(ch)

	for _,v := range list{
		ch <- struct{}{}
		v := v
		eg.Go(func() error {
			defer func() {
				<- ch
			}()

			select {
			case <- ctx.Done():
				fmt.Println("canceled:",v)
				return nil
			default:
				if v == "sdada"{
					return errors.New("the value is 123")
				}else if v == "123"{
					return errors.New("the value is hello")
				}
				return nil
			}
		})
	}

	if err := eg.Wait();err != nil {
		fmt.Println("err:",err)
	}
}
