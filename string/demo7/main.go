package main

import "fmt"

func main() {
	s1 := []string{"a","basd","1"}
	s2 := []string{"123","asdadad","sqweq","zcz"}

	a1 := []byte{104,101,108,108,111}
	fmt.Println(string(a1))

	copy(a1,"wo")

	fmt.Println("a1::",string(a1))

	copy(s2,s1)

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println([]byte("hello"))
}
