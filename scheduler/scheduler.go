/*
	YO WE ARE GONNA NEED WGs FOR THIS WHEN LOOPING THROUGH EACH EMPLOYEE AND SCHEDULING THEM
*/

// some basic logic here:
// 1. Need to loop through each day of the week, add the worker no more than 3 times to the WEEKLY schedule.
// 2. Loop through the range of days in the MONTHLY calendar, add those Weeklies to it.
package scheduler

import (
	"fmt"

	"github.com/brendisurfs/scheduler-app/workers"
)

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
)

var Weekdays = []int{Monday, Tuesday, Wednesday, Thursday, Friday}
var Calendar = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

type Schedule struct {
	Week int
	WS   map[int]Daily
}

type Daily struct {
	Day []workers.Worker
}

var WorkCalendar map[int]Daily
var DaySchedule Daily

// ScheduleWorker - places the worker on a slot, increments their placement by 1.
func scheduleWorker(w *workers.Worker) {
	// loop through each day in the calendar, assign worker day 0-4, UNLESS they have off.
	// 1. place the worker on their starting slot.
	// 2. check their days off, is it a day that is being scheduled? Dont put them on that day.
	for range Weekdays {
		if w.IsOff() {
			w.WorkDays = append(w.WorkDays, nil)
		}
	}
	w.WorkDays = append(w.WorkDays, *w)
	//
	// append the worker to the appropriate workdays.
	DaySchedule.Day = append(DaySchedule.Day, *w)
	// Finally, increment worker values (CurrentSlot, nextSlot)
	w.IncrementWorkerDay()

}

// AddToSchedule - adds each person to the global schedule.
func AddToCalendar(w *workers.Worker) {

	// ws := make(map[int]Daily)

	scheduleWorker(w)

	// let the user know the scheduling is done
	fmt.Println("Scheduling is done for: ", w.Name)
}
