package queue

import (
	"reflect"
)

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type queue struct {
	value []interface{}
	len   int
	types []string
}

func New(size int) Queue {
	return &queue{
		value: make([]interface{}, 0, size),
		len:   0,
		types: make([]string, 3),
	}
}

func (q *queue) Push(key interface{}) {
	q.value = append(q.value, key)

	typeof := reflect.TypeOf(key)
	if len(q.types) == 0 {
		q.types[0] = typeof.String()
		q.len++
	}

}

func (q *queue) Pop() interface{} {
	first := q.value[0]   //get the first element
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
	return len(q.value)
}

func (q *queue) Keys() []interface{} {
	return q.value
}
