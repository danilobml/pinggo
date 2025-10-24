package output

import (
	"fmt"

	"github.com/danilobml/pinggo/internal/models"
)

func PrintSummary(results models.SummaryResponse) {
	fmt.Println("**PINGGO RESULTS SUMMARY:**")

	fmt.Println()

	fmt.Printf("Total successful pings: %d\n", results.TotalSuccesses)
	fmt.Printf("Total slow pings: %d\n", results.TotalSlow)
	fmt.Printf("Total failed pings: %d\n", results.TotalErrors)

	fmt.Println()

	fmt.Printf("Average latency (succesfull calls): %v\n", results.AverageLatency)

	if results.TotalSuccesses > 0 {
		fmt.Println()
		fmt.Println("Successfuly pinged Urls:")
		for _, url := range results.SuccessUrls {
			fmt.Printf("%s\n", url)
		}
	}

	if results.TotalSlow > 0 {
		fmt.Println()
		fmt.Println("Slow Urls:")
		for _, url := range results.SlowUrls {
			fmt.Printf("%s\n", url)
		}
	}

	if results.TotalErrors > 0 {
		fmt.Println()
		fmt.Println("Failed URLs:")
		for _, url := range results.FailedUrls {
			fmt.Printf("%s\n", url)
		}
	}
}
