package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/thoriqadillah/screening/http/model/graduees"
	"github.com/thoriqadillah/screening/http/model/request"
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

func (g *GraduationAPI) GetGraduees(req *request.Request, wg *sync.WaitGroup) error {
	defer wg.Done()

	res, err := g.client.Get(req.URL())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&g.data); err != nil {
		return err
	}

	req.Save(&g.data)

	return nil
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
