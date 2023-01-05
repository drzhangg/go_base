package main

type User struct {
	Name string `json:"name"`
	Age string `json:"age"`
}

func main() {
	m := make(map[string]interface{})
	m["name"] = "jerry"
	m["age"] = 26
	m["address"] = "shanghai"


	//u := &User{}


}
