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

	gocron.Every(3).Seconds().Do(DdcOrderGenerateOrder)
	// gocron.Every(30).Seconds().Do(PtgRemoveExpiredInvitation)
	// gocron.Every(30).Seconds().Do(PtgRemoveExpiredCheckout)

	gocron.Every(1).Minutes().Do(CancelCustomDelivery)

	gocron.Every(5).Minutes().Do(VoucherRemoveReserveAndHistory)
	gocron.Every(1).Minute().Do(VoucherCreateHistory)
	gocron.Every(1).Minute().Do(VoucherCloseExpired)
	gocron.Every(5).Minutes().Do(PromoCloseExpired)

	gocron.Every(10).Minutes().Do(CancelOrder11)
	gocron.Every(10).Minutes().Do(CancelOrder12)
	gocron.Every(10).Minutes().Do(CancelOrder14)
	gocron.Every(10).Minutes().Do(CancelOrder16)

	// Runs foo and bar
	gocron.RunAll()

	// Starts the schedule as per normal
	<-gocron.Start()
}
