package main

import "fmt"

var (
	Label = map[string]string{"critical":"normal","gpu":"gpu"}
)

func main() {
	m := make(map[string]string)
	labels := Label
	labels["name"] = "jerry"

	sm := map[string]struct{}{"age":{}}
	_, ok := sm["age1"]
	fmt.Println("ok::",ok)

	m["jerry"] = "shanghai "

	fmt.Println("mmmm::", len(m["zhang"]))

	v,ok := m["jerry1"]
	fmt.Println(v,ok)

	fmt.Println("labels::",labels)

	for k,v := range m{
		fmt.Println(k,v)
	}
}
