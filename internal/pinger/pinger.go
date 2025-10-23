package pinger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/danilobml/pinggo/internal/errs"
)

type PingUrlResponse struct {
	StatusCode int
	Error bool
	Latency time.Duration
}

func PingFileUrls(filePath string) error {
	urls, err := getUrlsFromFile(filePath)
	if err != nil {
		return err
	}

	for _, url := range urls {
		pingResponse, _ := pingUrl(url)
		fmt.Printf("%+v\n", pingResponse)
	}


	return nil
}

func pingUrl(url string) (PingUrlResponse, error) {
	start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        return PingUrlResponse{
			Error: true,
		}, errs.ErrPingFailed
    }
	latency := time.Since(start)
    defer resp.Body.Close()

	pingResp := PingUrlResponse{
		StatusCode: resp.StatusCode,
		Error: false,
		Latency: latency,
	}

	return pingResp, nil
}
