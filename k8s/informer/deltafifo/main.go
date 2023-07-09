package main

import (
	"fmt"
	"k8s.io/client-go/tools/cache"
)

type pod struct {
	name  string
	value int
}

func newpod(name string, value int) pod {
	return pod{
		name:  name,
		value: value,
	}
}

type KeyFunc func(obj interface{}) (string, error)

func option(obj interface{}) (string, error) {
	result := obj.(pod)
	return result.name, nil
}

func main() {

	p1 := pod{
		name:  "p1",
		value: 1,
	}
	p2 := pod{
		name:  "p2",
		value: 2,
	}
	p3 := pod{
		name:  "p3",
		value: 3,
	}

	c := cache.NewDeltaFIFOWithOptions(cache.DeltaFIFOOptions{
		KeyFunction:           option,
	})

	c.Add(p1)
	c.Add(p2)
	c.Add(p3)

	p1.value = 4
	c.Update(p1)


	c.Pop(func(obj interface{}) error {
		for _, delta := range obj.(cache.Deltas){
			fmt.Println(delta.Type,":",delta.Object.(pod).name)

			switch delta.Type {
			case cache.Added:
				fmt.Println("add :",delta.Object.(pod).name)
			case cache.Updated:
				fmt.Println("update: ",delta.Object.(pod).name)
			case cache.Deleted:
				fmt.Println("delete: ",delta.Object.(pod).name)
			}
		}
		return nil
	})

}
