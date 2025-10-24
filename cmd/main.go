package main

import (
	"flag"

	"github.com/danilobml/pinggo/internal/pinger"
)

func main() {
	filePath := flag.String("from-file", "", "Pass a .txt file path to read the URLs to ping from it.")

	noCli := flag.Bool("no-cli", false, "Set this boolean flag, if you don't want a summary to be printed to the CLI (default false).")
	
	printJson := flag.Bool("json", false, "Set this boolean flag to create a json file with results.")
	
	flag.Parse()

	options := pinger.Options{
		PrintSummary: !*noCli,
		PrintJson: *printJson,
	}

	pinger.PingFileUrls(*filePath, options)
}
