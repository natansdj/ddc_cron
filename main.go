package main

import "github.com/jasonlvhit/gocron"

func main() {
	// remove, clear and next_run
	//_, timeStr := gocron.NextRun()
	//fmt.Println(timeStr)

	//gocron.Remove(task)
	//gocron.Clear()

	// function Start start all the pending jobs
	//<- gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	//s := gocron.NewScheduler()
	//s.Every(3).Seconds().Do(task)
	//<- s.Start()

	gocron.Every(1).Seconds().Do(DdcOrderGenerateOrder)
	// gocron.Every(30).Seconds().Do(RemoveExpiredInvitation)
	// gocron.Every(30).Seconds().Do(RemoveExpiredCheckout)

	// Runs foo and bar
	gocron.RunAll()

	// Starts the schedule as per normal
	<-gocron.Start()
}
