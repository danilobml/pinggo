package main

import (
	"flag"

	"github.com/danilobml/pinggo/internal/pinger"
)

func main() {
	filePath := flag.String("from-file", "", "Pass a .txt file path to read the URLs to ping from it.")

	noSummary := flag.Bool("no-summary", false, "Set this boolean flag, if you don't want a summary to be printed to the CLI (default false).")
	
	flag.Parse()

	options := pinger.Options{}

	options.Summary = !*noSummary

	pinger.PingFileUrls(*filePath, options)
}
