package main

import (
	"flag"

	"github.com/danilobml/pinggo/internal/pinger"
)

func main() {
	filePath := flag.String("from-file", "", "Pass a .txt file path to read the URLs to ping from it.")
	flag.Parse()

	pinger.PingFileUrls(*filePath)
}
