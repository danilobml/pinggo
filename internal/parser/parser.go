package parser

import (
	"bufio"
	"os"
	"path"

	"github.com/danilobml/pinggo/internal/errs"
	"github.com/danilobml/pinggo/internal/helpers"
)

func GetUrlsFromFile(filePath string) ([]string, error) {
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
		if helpers.IsValidURL(url) {
			urls = append(urls, url)
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return urls, nil
}
