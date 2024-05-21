package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func getWeek() int {
	baseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("CST", 8*3600))
	week := int(time.Now().Sub(baseDate).Hours()/24/7) + 1
	return week
}

func main() {

	fmt.Println(getWeek())

	// create a scheduler
	s:= gocron.NewScheduler(time.Local)
	//if err != nil {
	//	handle error
	//}

	// add a job to the scheduler
	//

	s.Cron("*/1 * * * *").Do(func() {
		fmt.Println("asdad",time.Now().String())
	})

	s.StartAsync()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute*10):
	}

	// when you're done, shut it down
	//err = s.Shutdown()
	//if err != nil {
	//	// handle error
	//}
}