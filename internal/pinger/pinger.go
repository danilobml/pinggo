package pinger

import (
	"context"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/danilobml/pinggo/internal/analizer"
	"github.com/danilobml/pinggo/internal/models"
	"github.com/danilobml/pinggo/internal/output"
	"github.com/danilobml/pinggo/internal/parser"
)

const (
	genTimeout     = 15 * time.Second
	requestTimeout = 5 * time.Second
)

type Options struct {
	PrintSummary bool
	PrintJson    bool
	Concurrency  int
}

func PingFileUrls(filePath string, options Options) error {
	if options.Concurrency <= 0 {
		options.Concurrency = runtime.NumCPU()
	}

	urls, err := parser.GetUrlsFromFile(filePath)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), genTimeout)
	defer cancel()

	var wg sync.WaitGroup
	resultsChan := make(chan models.Result, options.Concurrency)
	jobsChan := make(chan string, options.Concurrency) 

	pingerResponse := models.PingerResponse{}

	client := &http.Client{
		Timeout: requestTimeout,
		Transport: &http.Transport{
			MaxConnsPerHost: options.Concurrency,
		},
	}

	// worker pool
	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobsChan {
				result := pingUrl(ctx, client, url)
				select {
				case resultsChan <- result:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		defer close(jobsChan)
		for _, url := range urls {
			select {
			case jobsChan <- url:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for {
		select {
		case response, ok := <-resultsChan:
			if !ok {
				summary := analizer.GenerateSummary(pingerResponse)

				if options.PrintSummary {
					output.PrintTable(summary)
				}

				if options.PrintJson {
					output.PrintJson(summary)
				}
				return nil
			}
			pingerResponse = append(pingerResponse, response)
		case <-ctx.Done():
			summary := analizer.GenerateSummary(pingerResponse)

			if options.PrintSummary {
				output.PrintTable(summary)
			}

			if options.PrintJson {
				output.PrintJson(summary)
			}
			return ctx.Err()
		}
	}
}

func pingUrl(ctx context.Context, client *http.Client, url string) models.Result {
	reqCtx, cancel := context.WithTimeout(ctx, client.Timeout)
	defer cancel()

	start := time.Now()
	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		return models.Result{
			Url:   url,
			Error: true,
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.Result{Url: url, Error: true}
	}
	latency := time.Since(start)
	defer resp.Body.Close()

	pingResp := models.Result{
		Url:        url,
		StatusCode: resp.StatusCode,
		Error:      false,
		Latency:    latency,
	}

	return pingResp
}
