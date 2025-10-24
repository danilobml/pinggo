package models

import "time"

type PingerResponse = []Result

type Result struct {
	Url        string
	StatusCode int
	Error      bool
	Latency    time.Duration
}

type SummaryResponse struct {
	TotalSuccesses int
	TotalSlow      int
	TotalErrors    int
	AverageLatency time.Duration
	SuccessUrls    []string
	SlowUrls       []string
	FailedUrls     []string
}
