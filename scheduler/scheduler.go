package scheduler

import (
	"github.com/brendisurfs/scheduler-app/workers"
)

// ScheduleWorker - places the worker on a slot, increments their placement by 1.
func ScheduleWorker(w *workers.Worker) {
	// loop through each day in the calendar, assign worker day 0-4, UNLESS they have off.
	w.IncrementWorkerDay()
}
