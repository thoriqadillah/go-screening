package main

import (
	"github.com/thoriqadillah/screening/cmd"
	"github.com/thoriqadillah/screening/enitity"
	"github.com/thoriqadillah/screening/http/api"
	"github.com/thoriqadillah/screening/http/service"
	"github.com/thoriqadillah/screening/lib/array"
)

func main() {
	numbers := []int{2, 3, 1, 5, 3}
	width := len(numbers) * 2
	height := array.GetMaxNumber(numbers[:])

	canvas := enitity.NewCanvas(width, height).Draw()
	array.Sort(numbers[:], func(a, b int) bool {
		canvas.DrawChart(numbers)
		canvas.Display()

		return b > a //descending
		// return a > b //ascending
	})

	var concurrent_limit int
	var output string

	cmd.Parse(&concurrent_limit, &output)

	const URL = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"
	api := api.NewGraduationAPI(URL)
	graduation := service.NewGraduation(api)

	graduation.ToCSV(output, concurrent_limit, "2013", "2000", "2001")
}
