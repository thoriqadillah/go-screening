package request

import (
	"github.com/thoriqadillah/screening/http/model/graduees"
)

type Request struct {
	url    string
	name   string
	result *graduees.Data
}

func New(url string, name string) Request {
	return Request{
		url:    url,
		name:   name,
		result: &graduees.Data{},
	}
}

func (r *Request) URL() string {
	return r.url
}

func (r *Request) Name() string {
	return r.name
}

func (r *Request) Result() *graduees.Data {
	return r.result
}

func (r *Request) Save(res *graduees.Data) {
	r.result = res
}
