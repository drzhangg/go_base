package main

import "fmt"

type Data struct {
	User []*User
}

type User struct {
	name   string
	mobile []*phone
	age    int
}

type phone struct {
	key string
}

func main() {
	d := &Data{}
	fmt.Println(getUser(d))

}

func getUser(data *Data) bool {
	//if data.User == nil {
	//	return false
	//}
	us := data.User

	//keys := make(map[string]struct{})
	keys := map[string]struct{}{"110":{},"120":{}}
	for _, u := range us{
		for _,p := range u.mobile{
			if _,ok:= keys[p.key];ok{
				fmt.Println(keys)
				return true
			}
		}
	}
	return false
}
