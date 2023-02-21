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

	URL := g.api.URL()

	workers := worker.NewWorker(concurrent_limit)
	workers.Run()

	var wg sync.WaitGroup

	for _, year := range years {
		temp := URL
		query := "&q=" + year
		temp += query
		g.api.UpdateURL(temp)

		wg.Add(2)
		req := request.New(g.api.URL(), &wg)
		workers.Add(func() {
			g.api.GetGraduees(&req)
		})

		workers.Add(func() {
			p := path + "/" + year + ext
			g.writeCSV(p, &req)
		})

		g.api.UpdateURL(URL)
	}

	wg.Wait()

	return nil
}

func (g *GraduationService) writeCSV(name string, req *request.Request) {
	defer req.Done()

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	csvwriter := csv.NewWriter(file)
	grads := <-req.Result()

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
