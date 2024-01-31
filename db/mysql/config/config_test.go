package config

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	db,err:= InitMysql()

	m := make(map[string]interface{})
	err = db.Table("account").Find(&m).Error
	if err != nil {
		fmt.Println("err:",err)
	}

	fmt.Println(m)
}
