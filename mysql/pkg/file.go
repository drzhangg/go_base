package pkg

import (
	"bufio"
	"log"
	"os"
)

//var FileChan = make(chan string)

func GetFile(ch chan string) {

	path := "/Users/drzhang/demo/go/go_base/mysql/insert/data.sql"

	f, err := os.Open(path)
	if err != nil {
		log.Println("error:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ch <- scanner.Text()
	}

	defer close(ch)

}
