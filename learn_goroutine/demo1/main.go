package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		 "home": "415.333.3333",
		 "cell": "415.555.5555",
		 }

	 // 将这个映射序列化到

	 data,err := json.MarshalIndent(c,"","  ")
	 if err != nil {
		 fmt.Println("err:",err)
		 return
	 }

	 fmt.Println(string(data))


}
