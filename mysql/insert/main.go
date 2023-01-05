package main

import (
	"fmt"
	"go_base/mysql/config"
	"go_base/mysql/pkg"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func main() {
	db, err := config.InitMysql()
	if err != nil {
		fmt.Println(err)
	}

	inch := make(chan string)

	go pkg.GetFile(inch)

	gnum := 10

	eg := errgroup.Group{}
	for i := 0; i < gnum; i++ {

		eg.Go(func() error {

			for v := range inch {
				err := insert(db, v)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}

	if err := eg.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}

func insert(db *gorm.DB, sql string) error {

	return db.Exec(sql).Error
}

func generateSql(num int, ch chan string) {
	sql := "INSERT INTO `books` VALUES ('%d', 'Java项目实战教程+Java程序设计与项目实训教程ccc %d', '', '72.00', '72.50', '姜华刘闯', '清华大学出版社', ' 2012-09-01', 'PHP', '2', '%d');"

	rand.Seed(time.Now().UnixNano())

	randNum := rand.Intn(7)
	ch <- fmt.Sprintf(sql, num, num, randNum)
	defer close(ch)
}
