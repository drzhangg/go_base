package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var Db *gorm.DB
func InitMysql() (*gorm.DB,error) {
	dsn := "root:root@tcp(150.158.87.137:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("connect mysql failed:%v", err)
		return nil,err
	}

	//Db = db
	return db, nil
}
