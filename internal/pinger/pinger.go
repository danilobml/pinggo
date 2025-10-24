package pinger

import (
	"net/http"
	"time"

	"github.com/danilobml/pinggo/internal/analizer"
	"github.com/danilobml/pinggo/internal/errs"
	"github.com/danilobml/pinggo/internal/models"
	"github.com/danilobml/pinggo/internal/output"
	"github.com/danilobml/pinggo/internal/parser"
)

type Options struct {
	Summary bool
}

func PingFileUrls(filePath string, options Options) error {
	urls, err := parser.GetUrlsFromFile(filePath)
	if err != nil {
		return err
	}

	pingerResponse := models.PingerResponse{}

	for _, url := range urls {
		pingResponse, _ := pingUrl(url)
		pingerResponse = append(pingerResponse, pingResponse)
	}

	if options.Summary {
		summary := analizer.GenerateSummary(pingerResponse)
		output.PrintSummary(summary)
	}

	return nil
}

func pingUrl(url string) (models.Result, error) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return models.Result{
			Url:   url,
			Error: true,
		}, errs.ErrPingFailed
	}
	latency := time.Since(start)
	defer resp.Body.Close()

	pingResp := models.Result{
		Url:        url,
		StatusCode: resp.StatusCode,
		Error:      false,
		Latency:    latency,
	}

	return pingResp, nil
}
