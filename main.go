package main

import (
	"fmt"
	"os"

	"github.com/thoriqadillah/screening/dir"
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

	sourcepath := "/home/thoriqadillah/Development/Go/src/hello"
	source := dir.NewDirectory(sourcepath)
	sourcefiles := source.Scan()
	// if err := persistent.Save("source.tmp", f1); err != nil {
	// 	panic(err)
	// }

	fmt.Println("========")

	targetpath := "/home/thoriqadillah/Development/Go/src/hello-copy"
	target := dir.NewDirectory(targetpath)
	targetfiles := target.Scan()
	// if err := persistent.Save("target.tmp", f2); err != nil {
	// 	panic(err)
	// }

	// for _, file := range sourcefiles {
	// 	_, err := os.Stat(p2 + file.Name)
	// 	if err != nil {
	// 		fmt.Println(file.Name + " DELETED")
	// 	}
	// }

	for _, sourcefile := range sourcefiles {
		_, err := os.Stat(targetpath + sourcefile.Name)

		if err != nil && sourcefile.CreatedAt.Equal(sourcefile.ModifiedAt) {
			fmt.Println(sourcepath + sourcefile.Name + " NEW")
		}
	}

	for _, targetfile := range targetfiles {
		_, err := os.Stat(sourcepath + targetfile.Name)

		if err != nil && !targetfile.CreatedAt.Equal(targetfile.ModifiedAt) {
			fmt.Println(sourcepath + targetfile.Name + " DELETED")
		}
	}

	// var concurrent_limit int
	// var output string

	// cmd.Parse(&concurrent_limit, &output)

	// const URL = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"
	// api := api.NewGraduationAPI(URL)
	// graduation := service.NewGraduation(api)

	// graduation.ToCSV(output, concurrent_limit, "2013", "2000", "2001")
}
