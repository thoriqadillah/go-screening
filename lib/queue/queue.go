package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
	In() []interface{}
}

type queue struct {
	in    []interface{}
	value []interface{}
	size  int
	len   int
}

func New(size int) Queue {
	return &queue{
		in:    make([]interface{}, 0, size),
		value: make([]interface{}, 0, size),
		size:  size,
		len:   0,
	}
}

func (q *queue) Push(key interface{}) {
	if q.Contains(key) {
		return
	}

	q.value = append(q.value, key)
	q.len++

	if q.len <= q.size {
		q.in = append(q.in, key)
	}
}

func (q *queue) Pop() interface{} {
	first := q.in[0] //get the first element

	q.in = q.in[1:]
	q.value = q.value[1:] //remove the first element
	q.len--

	return first
}

func (q *queue) Contains(key interface{}) bool {
	for _, v := range q.value {
		if v == key {
			return true
		}
	}

	return false
}

func (q *queue) Len() int {
	return len(q.in)
}

func (q *queue) Keys() []interface{} {
	return q.value
}

func (q *queue) In() []interface{} {
	return q.in
}
