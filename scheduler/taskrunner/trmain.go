package taskrunner

import "time"

/*
timer
setup
start{trigger->task->runner}
*/

type Worker struct {
	ticker *time.Ticker
	runner *Runner

}

func NewWork(interval time.Duration,r *Runner) *Worker  {

	return &Worker{
		ticker:time.NewTicker(interval*time.Second),
		runner:r,
	}
}

func (w *Worker)startWork()  {
	for  {
		select {
		case <- w.ticker.C:
			go w.runner.StartAll()
		}

	}
}
func Start()  {
	// Start video file cleaning
	r := NewRunner(3,true,VideoClearDispatcher,VideoClearExecutor)
	w := NewWork(3,r)
	go w.startWork()
}