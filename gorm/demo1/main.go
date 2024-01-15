package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	Name string `gorm:"column:name;not null" json:"name"`
	Ip   string `gorm:"column:ip;not null;unique" json:"ip"`
	Host string `gorm:"column:hostip" json:"host"`
	Desc string `gorm:"column:desc" json:"desc"`
}

func (*Server) TableName() string {
	return "server"
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(" gorm.Open err:",err)
		return
	}

	name := []string{}

	//m := make(map[string]interface{})
	err = db.Table("students").Select("name").Scan(&name).Error
	if err != nil {
		fmt.Println("select err:",err)
		return
	}

	fmt.Println(name)


	//var (
	//	hostList []string
	//)

	//tx := db.GORM.Table("server").Select("hostip").Find(&hostList)


}
