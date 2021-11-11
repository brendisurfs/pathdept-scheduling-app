/*
	YO WE ARE GONNA NEED WGs FOR THIS WHEN LOOPING THROUGH EACH EMPLOYEE AND SCHEDULING THEM
*/

// some basic logic here:
// 1. Need to loop through each day of the week, add the worker no more than 3 times to the WEEKLY schedule.
// 2. Loop through the range of days in the MONTHLY calendar, add those Weeklies to it.
package scheduler

import (
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
	// 2. check their days off, is it a day that is being scheduled? Dont put them on that day.
	// NOTE: i am just ranging through the days that can work, but only 0-4. FIXXXXX MEEE
	for _, v := range Weekdays {
		if w.IsOff() {
			w.WorkDays = append(w.WorkDays, nil)
		}
		w.WorkDays = append(w.WorkDays, v)
	}
	DaySchedule.Day = append(DaySchedule.Day, *w)
	//
	// append the worker to the appropriate workdays.
	// Finally, increment worker values (CurrentSlot, nextSlot)
	w.IncrementWorkerDay()

}

// AddToSchedule - adds each person to the global schedule.
func AddToCalendar(workers *[]workers.Worker) {

	WorkCalendar = make(map[int]Daily)

	// for each week

	// for each worker, schedule them and add them to the calendar.
	for _, w := range *workers {
		scheduleWorker(&w)
	}
	// ... and add a daily schedule to each day

	// WorkCalendar[i] = DaySchedule

	// let the user know the scheduling is done
}
