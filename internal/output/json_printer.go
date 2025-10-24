package output

import (
	"encoding/json"
	"os"

	"github.com/danilobml/pinggo/internal/models"
)

func PrintJson(results models.SummaryResponse) error {
	file, err := os.Create("./results.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	formattedResults := map[string]any{
		"total_successes":             results.TotalSuccesses,
		"slow_requests":               results.TotalSlow,
		"failed_requests":             results.TotalErrors,
		"average_latency_microseconds": results.AverageLatency.Microseconds(),
		"successful_urls":             results.SuccessUrls,
		"slow_urls":                   results.SlowUrls,
		"failed_urls":                 results.FailedUrls,
	}

	if err := encoder.Encode(formattedResults); err != nil {
		return err
	}

	return nil
}
