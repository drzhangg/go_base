package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "/Users/drzhang/test/yaml/test-sts.yaml")
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Printf("combined out:\n%s\n", string(out))
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
	//fmt.Printf("combined out:\n%s\n", string(out))

	output,err := cmd.Output()
	if err != nil {
			fmt.Printf("cmd out err:%v\n", err)
			return
	}

	fmt.Println("out::",string(output))
}
