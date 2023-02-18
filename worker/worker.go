package worker

import "fmt"

type worker struct {
	total int
	jobs  chan func()
}

func NewWorker(total int) *worker {
	return &worker{
		total: total,
		jobs:  make(chan func()),
	}
}

func (w *worker) Run() {
	for i := 0; i < w.total; i++ {
		go func(id int) {
			fmt.Printf("WORKER %d IS WORKING\n", id)
			for job := range w.jobs {

				job()
				fmt.Printf("WORKER %d IS DONE\n", id)
			}
		}(i)
	}
}

func (w *worker) Add(job func()) {
	w.jobs <- job
}
