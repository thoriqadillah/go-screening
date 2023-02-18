package main

import (
	"fmt"

	"github.com/thoriqadillah/screening/dir"
	"github.com/thoriqadillah/screening/dir/persistent"
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

	source := dir.NewDirectory("/home/thoriqadillah/Development/Go/src/go-screening")
	f1 := source.Scan()
	if err := persistent.Save("source.tmp", f1); err != nil {
		panic(err)
	}

	fmt.Println("========")

	target := dir.NewDirectory("/home/thoriqadillah/Development/Go/src/go-screening-copy")
	f2 := target.Scan()
	if err := persistent.Save("target.tmp", f2); err != nil {
		panic(err)
	}

	// var concurrent_limit int
	// var output string

	// cmd.Parse(&concurrent_limit, &output)

	// const URL = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"
	// api := api.NewGraduationAPI(URL)
	// graduation := service.NewGraduation(api)

	// graduation.ToCSV(output, concurrent_limit, "2013", "2000", "2001")
}
