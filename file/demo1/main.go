package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	ch := make(chan string)
	defer close(ch)

	num := 5
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		go func() {
			for v := range ch {
				fmt.Println(v)

				file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					fmt.Println("文件打开失败", err)
				}
				//及时关闭file句柄
				defer file.Close()

				write := bufio.NewWriter(file)

				write.WriteString(v + "\r\n")

				//Flush将缓存的文件真正写入到文件中
				write.Flush()
			}
		}()
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go generateSql(i, ch, &wg)
	}

	wg.Wait()
}

func generateSql(num int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	sql := "INSERT INTO `books` VALUES ('%d', 'Java项目实战教程+Java程序设计与项目实训教程ccc %d', '', '72.00', '72.50', '姜华刘闯', '清华大学出版社', ' 2012-09-01', 'PHP', '2', '%d');"

	randNum := rand.Intn(7)
	ch <- fmt.Sprintf(sql, num, num, randNum)
}
