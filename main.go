package main

import (
	"github.com/thoriqadillah/screening/cmd"
	"github.com/thoriqadillah/screening/http/api"
	"github.com/thoriqadillah/screening/http/service"
	"github.com/thoriqadillah/screening/lib/array"
	"github.com/thoriqadillah/screening/renderer"
)

func main() {
	numbers := []int{2, 3, 1, 5, 3}
	width := len(numbers) * 2
	height := array.Max(numbers[:]) + 1

	canvas := renderer.NewCanvas(width, height).Draw()
	array.Sort(numbers[:], func(a, b int) bool {
		canvas.DrawChart(numbers)
		canvas.Display()

		return b > a //descending
		// return a > b //ascending
	})

	// sourcepath := "/home/thoriqadillah/Development/Go/src/go-screening"
	// source := dir.NewDirectory(sourcepath)

	// targetpath := "/home/thoriqadillah/Development/Go/src/go-screening-copy"
	// target := dir.NewDirectory(targetpath)

	// source.Compare(&target)

	var concurrent_limit int
	var output string

	cmd.Parse(&concurrent_limit, &output)

	const URL = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"
	api := api.NewGraduationAPI(URL)
	graduation := service.NewGraduation(api)

	graduation.ToCSV(output, concurrent_limit, "2000", "2001", "2002", "2003", "2004", "2005", "2006", "2007", "2008", "2009", "2010", "2011", "2012", "2013", "2014")
}
