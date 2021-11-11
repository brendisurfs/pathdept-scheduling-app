/*
	YO WE ARE GONNA NEED WGs FOR THIS WHEN LOOPING THROUGH EACH EMPLOYEE AND SCHEDULING THEM
*/
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

// WorkSchedule - a global schedule that holds information about each person on which day.
var WorkSchedule map[int]Day

type Schedule struct {
	Week int
}

type Day struct {
	Workers []workers.Worker
}

// ScheduleWorker - places the worker on a slot, increments their placement by 1.
func ScheduleWorker(w *workers.Worker) {
	// loop through each day in the calendar, assign worker day 0-4, UNLESS they have off.
	newDay := Day{}
	for i, v := range workers.Calendar {
		// 1. place the worker on their starting day.
		// 2. check their days off, is it a day that is being scheduled? Dont put them on that day.
		for _, d := range w.DaysOff {
			if d == v {
				w.WorkDays = append(w.WorkDays, nil)
			}
		}
		w.WorkDays = append(w.WorkDays, i)
	}
	//
	// append the worker to the appropriate workdays.
	newDay.Workers = append(newDay.Workers, *w)

	// Finally, increment the workers day.
	w.IncrementWorkerDay()
}
