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

	gocron.Every(5).Minutes().Do(CancelOrder, "1", "0")
	// gocron.Every(5).Minutes().Do(CancelOrder, "1", "1")
	//gocron.Every(5).Minutes().Do(CancelOrder, "1", "2")
	//gocron.Every(5).Minutes().Do(CancelOrder, "1", "4")
	//gocron.Every(5).Minutes().Do(CancelOrder, "1", "6")

	gocron.Every(5).Minutes().Do(VoucherRemoveReserveAndHistory)
	gocron.Every(5).Minutes().Do(VoucherCreateHistory)
	gocron.Every(5).Minutes().Do(VoucherCloseExpired)
	gocron.Every(5).Minutes().Do(PromoCloseExpired)

	gocron.Every(5).Minutes().Do(CancelCustomDelivery)

	gocron.Every(5).Minutes().Do(CheckMemberExpired)
	gocron.Every(5).Minutes().Do(CheckMemberExpiredPoint)
	gocron.Every(5).Minutes().Do(UnlockMember, "1")
	//gocron.Every(1).Day().Do(EmailStepExpireDate, "-3")

	gocron.Every(5).Minutes().Do(PtgRemoveExpiredInvitation)
	gocron.Every(5).Minutes().Do(PtgRemoveExpiredCheckout)
	gocron.Every(30).Minutes().Do(PtgLockInitiator)

	// Runs foo and bar
	gocron.RunAll()

	// Starts the schedule as per normal
	<-gocron.Start()
}
