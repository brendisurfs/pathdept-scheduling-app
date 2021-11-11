package main

import (
	"fmt"

	"github.com/brendisurfs/scheduler-app/scheduler"
	"github.com/brendisurfs/scheduler-app/workers"
)

const MAX_WORKERS_OFF = 2

func main() {

	dm := workers.Worker{
		Name:         "dm",
		StartingSlot: 0,
		CurrentSlot:  2,
		GlobalSlot:   2,
		DaysOff:      []int{2, 4, 6, 9, 11, 15, 18},
	}

	fmt.Println("dm days off: ", dm.DaysOff)
	fmt.Println(dm.IsOff())
	scheduler.ScheduleWorker(&dm)
	fmt.Println("work days: ", dm.WorkDays)
}
