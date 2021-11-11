/*
	YO WE ARE GONNA NEED WGs FOR THIS WHEN LOOPING THROUGH EACH EMPLOYEE AND SCHEDULING THEM
*/
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

type Schedule struct {
	Week int
	WS   map[int]Daily
}

type Daily struct {
	Day []workers.Worker
}

// WorkSchedule - a global schedule that holds information about each person on which day.
var WorkCalendar map[int]Daily
var DaySchedule Daily

// ScheduleWorker - places the worker on a slot, increments their placement by 1.
func scheduleWorker(i int, v int, w *workers.Worker) {
	// loop through each day in the calendar, assign worker day 0-4, UNLESS they have off.
	// 1. place the worker on their starting slot.
	// 2. check their days off, is it a day that is being scheduled? Dont put them on that day.
	for _, d := range w.DaysOff {
		if d == v {
			w.WorkDays = append(w.WorkDays, nil)
		}
	}
	w.WorkDays = append(w.WorkDays, i)
	//
	// append the worker to the appropriate workdays.
	DaySchedule.Day = append(DaySchedule.Day, *w)
	// Finally, increment worker values (CurrentSlot, nextSlot)
	w.IncrementWorkerDay()

}

// AddToSchedule - adds each person to the global schedule.
func AddToCalendar(w *workers.Worker) {

	ws := make(map[int]Daily)

	for i, v := range workers.Calendar {
		scheduleWorker(i, v, w)
		ws[i] = DaySchedule
	}
	WorkCalendar = ws
	// let the user know the scheduling is done
	fmt.Println("Scheduling is done for: ", w.Name)
}
