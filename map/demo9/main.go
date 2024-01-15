package main

func main() {

	m1 := make(map[string]string)
	m1["name"] = "tom"

	m2 := make(map[string]interface{})

	for k,v := range m1{
		m2[k] = v
	}
}
