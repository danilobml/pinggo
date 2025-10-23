package pinger

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/danilobml/pinggo/internal/errs"
)

func PingFiles(filePath string) error {
	urls, err := getUrlsFromFile(filePath)
	if err != nil {
		return err
	}

	for _, url := range urls {
		fmt.Println(url)
	}

	return nil
}

func getUrlsFromFile(filePath string) ([]string, error) {
	ext := path.Ext(filePath)
	if ext != ".txt" {
		return nil, errs.ErrInvalidInputFile
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	urls := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return urls, nil
}
