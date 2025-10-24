package output

import (
	"time"

	"github.com/danilobml/pinggo/internal/models"
)

const SlowThreshold = 999 * time.Millisecond

func FormatSummary(results models.PingerResponse) models.SummaryResponse {
	summary := models.SummaryResponse{}
	totalLatency := time.Duration(0)

	for _, result := range results {
		if result.Error {
			summary.TotalErrors++
			summary.FailedUrls = append(summary.FailedUrls, result.Url)
		} else {
			summary.TotalSuccesses++
			totalLatency += result.Latency
			summary.SuccessUrls = append(summary.SuccessUrls, result.Url)
			if result.Latency > SlowThreshold {
				summary.SlowUrls = append(summary.SlowUrls, result.Url)
				summary.TotalSlow++
			}
		}
	}

	if summary.TotalSuccesses > 0 {
		summary.AverageLatency = totalLatency / time.Duration(summary.TotalSuccesses)
	}

	return summary
}
