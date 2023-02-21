package worker

type queue struct {
	jobs chan func()
}

func NewQueue() queue {
	return queue{
		jobs: make(chan func()),
	}
}

func (q *queue) Add(job func()) {
	q.jobs <- job
}
