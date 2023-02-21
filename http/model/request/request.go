package request

import (
	"sync"

	"github.com/thoriqadillah/screening/http/model/graduees"
)

type Request struct {
	url    string
	result chan *graduees.Data
	wg     *sync.WaitGroup
}

func New(url string, wg *sync.WaitGroup) Request {
	return Request{
		url:    url,
		result: make(chan *graduees.Data),
		wg:     wg,
	}
}

func (r *Request) Save(res *graduees.Data) {
	r.result <- res
}

func (r *Request) Result() chan *graduees.Data {
	return r.result
}

func (r *Request) Done() {
	r.wg.Done()
}
