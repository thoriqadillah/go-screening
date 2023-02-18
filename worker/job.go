package worker

type Pool interface {
	Run()
	Add(task func())
}

type worker struct {
	total int
	queue chan func()
}

func NewWorker(total int) Pool {
	return &worker{
		total: total,
	}
}
func (w *worker) Run() {
	for i := 0; i < w.total; i++ {
		go func(id int) {
			for task := range w.queue {
				task()
			}
		}(i + 1)
	}
}

func (w *worker) Add(task func()) {
	w.queue <- task
}
