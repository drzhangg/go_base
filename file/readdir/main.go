package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	dir := "/Users/drzhang/demo/go/go_base"

	files, err := ioutil.ReadDir(dir)
	switch {
	case err == nil: // break
		fmt.Println("err is nil")
	case os.IsNotExist(err):
		fmt.Println("the dir is not exist")
	default:
		fmt.Println("case default")
	}

	for _,file := range files{

		if file.IsDir(){
			continue
		}


		fileExt := filepath.Ext(file.Name())
		fmt.Println(file.Name(),fileExt)
	}

	arr := []string{"asdasd","basdqweq","adaqweqwe","brwqd","cswqe"}
	sort.Strings(arr)
	fmt.Println(arr)
}
