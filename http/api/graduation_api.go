package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/thoriqadillah/screening/http/model/graduees"
)

type GraduationAPI struct {
	client *http.Client
	url    string
	data   graduees.Data
}

func NewGraduationAPI(url string) GraduationAPI {
	return GraduationAPI{
		client: &http.Client{},
		url:    url,
		data:   graduees.Data{},
	}
}

func (g *GraduationAPI) GetGraduees(ch chan *graduees.Data, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := g.client.Get(g.url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&g.data); err != nil {
		panic(err)
	}

	ch <- &g.data
}

func (g *GraduationAPI) GetData() *graduees.Data {
	return &g.data
}

func (g *GraduationAPI) UpdateURL(url string) {
	g.url = url
}

func (g *GraduationAPI) URL() string {
	return g.url
}
