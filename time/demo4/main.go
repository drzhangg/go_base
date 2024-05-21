package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getWeek())

}

func getWeek() int {
	baseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("CST", 8*3600))
	week := int(time.Now().Sub(baseDate).Hours()/24/7) + 1
	return week
}
