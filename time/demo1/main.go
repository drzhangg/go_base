package main

import (
	"fmt"
	"time"
)

func main() {
	unusedTime := "2022-01-30T14:15:31Z"

	loc,err := time.LoadLocation("Local")
	if err !=nil{
		fmt.Println(err)
	}

	locTime,err := time.ParseInLocation("2006-01-02T15:04:05Z",unusedTime,loc)
	if err !=nil{
		fmt.Println(err)
	}

	//fmt.Println(loc)
	//fmt.Println(locTime)

	now := time.Now()

	clearTime := locTime.Add(time.Hour * 24)
	sendTime := locTime.Add(time.Hour)

	fmt.Println(now)
	fmt.Println(clearTime)
	fmt.Println(sendTime)
	//fmt.Println(now.Before(clearTime))
	//fmt.Println(now.After(sendTime))
}
