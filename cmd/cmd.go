package cmd

import (
	"flag"
	"os"
)

func Parse(concurrent_limit *int, output *string) {
	help := flag.Bool("help", false, "Show help")

	flag.IntVar(concurrent_limit, "concurrent_limit", 2, "Limit the concurency")
	flag.StringVar(output, "output", "./storage", "The data location")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}
}
