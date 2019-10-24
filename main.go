package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"os/exec"
	"time"
)

func Output(str ...interface{}) {
	var output string
	output += time.Now().Format("01-01-2006 15:04:05.00") + " : "
	output += fmt.Sprintln(str...)

	fmt.Print(output)
}

func CurlGET(url string) {
	Output(url)

	curl := exec.Command("curl", url)
	out, err := curl.Output()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(out))
	}
}

func task() {
	fmt.Println("I am runnning task.")
}

func main() {
	// remove, clear and next_run
	_, timeStr := gocron.NextRun()
	fmt.Println(timeStr)

	//gocron.Remove(task)
	//gocron.Clear()

	// function Start start all the pending jobs
	//<- gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	//s := gocron.NewScheduler()
	//s.Every(3).Seconds().Do(task)
	//<- s.Start()

	gocron.Every(5).Seconds().Do(DdcOrderGenerateOrder)
	gocron.Every(10).Seconds().Do(RemoveExpiredInvitation)
	gocron.Every(10).Seconds().Do(RemoveExpiredCheckout)

	// Runs foo and bar
	gocron.RunAll()

	// Starts the schedule as per normal
	<-gocron.Start()
}
