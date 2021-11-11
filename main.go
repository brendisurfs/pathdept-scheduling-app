package main

import (
	"fmt"

	"github.com/brendisurfs/scheduler-app/scheduler"
	"github.com/brendisurfs/scheduler-app/workers"
)

func main() {

	dm := workers.Worker{
		Name:         "dr malicki",
		StartingSlot: 0,
		CurrentSlot:  2,
		GlobalSlot:   2,
		DaysOff:      []int{2, 3, 6, 7},
	}

	fmt.Println("dm days off: ", dm.DaysOff)
	fmt.Println(dm.IsOff())
	fmt.Println("Current Slot Before schedule worker: ", dm.CurrentSlot)
	scheduler.ScheduleWorker(&dm)
	fmt.Println("after schedule worker: ", dm.CurrentSlot)
	scheduler.ScheduleWorker(&dm)
	fmt.Println("after schedule worker: ", dm.CurrentSlot)
	scheduler.ScheduleWorker(&dm)
	fmt.Println("after schedule worker: ", dm.CurrentSlot)
}
