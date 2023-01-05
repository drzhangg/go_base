package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	strBytes, err := inputReader.ReadBytes('\n')
	if err != nil {
		fmt.Println("err：", err)
	}
	fmt.Println("strBytes：", string(strBytes))

}
