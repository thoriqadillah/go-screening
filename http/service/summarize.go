package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/thoriqadillah/screening/http/api"
	"github.com/thoriqadillah/screening/http/model/request"
	"github.com/thoriqadillah/screening/worker"
)

// var wg sync.WaitGroup
// var file = make(chan *os.File)
// var url = make(chan string)

type GraduationService struct {
	api api.GraduationAPI
}

func NewGraduation(api api.GraduationAPI) GraduationService {
	return GraduationService{
		api: api,
	}
}

func (g *GraduationService) ToCSV(path string, concurrent_limit int, years ...string) error {
	ext := ".csv"

	ch := make(chan *request.Request)
	workers := worker.NewWorker(concurrent_limit)
	workers.Run()

	var wg sync.WaitGroup

	for _, year := range years {
		url := g.api.URL()
		query := "&q=" + year
		url += query

		wg.Add(2)
		name := path + "/" + year + ext
		req := request.New(url, name)
		workers.Add(func() {
			if err := g.api.GetGraduees(&req, &wg); err != nil {
				panic(err)
			}
			ch <- &req
		})

		workers.Add(func() {
			g.writeCSV(ch, &wg)
		})
	}

	wg.Wait()

	return nil
}

func (g *GraduationService) writeCSV(ch chan *request.Request, wg *sync.WaitGroup) {
	defer wg.Done()

	req := <-ch

	file, err := os.Create(req.Name())
	if err != nil {
		panic(err)
	}

	csvwriter := csv.NewWriter(file)
	grads := req.Result()

	forfield := true

	rows := make([][]string, len(grads.Result.Records))
	lencol := reflect.TypeOf(grads.Result.Records[0]).NumField()
	fields := make([]string, lencol)
	for i, record := range grads.Result.Records {
		rows[i] = make([]string, lencol)

		val := reflect.Indirect(reflect.ValueOf(record))
		for j := 0; j < lencol; j++ {
			rows[i][j] = fmt.Sprint(val.Field(j).Interface())
			fields[j] = val.Type().Field(j).Name
		}
		if forfield {
			csvwriter.Write(fields)
			forfield = false
		}

		csvwriter.Write(rows[i])
	}

	csvwriter.Flush()
}
