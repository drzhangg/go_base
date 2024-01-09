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

	dsn := "root:root@tcp(150.158.87.137:3306)/daota?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	name := []string{}

	//m := make(map[string]interface{})
	err = db.Table("heros").Select("name").Scan(&name).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(name)


	//var (
	//	hostList []string
	//)

	//tx := db.GORM.Table("server").Select("hostip").Find(&hostList)


}
