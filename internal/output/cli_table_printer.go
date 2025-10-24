package output

import (
	"os"

	"github.com/danilobml/pinggo/internal/models"
	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintTable(results models.SummaryResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Ping Summary Report")

	t.AppendHeader(table.Row{"Metric", "Value"})
	t.AppendRows([]table.Row{
		{"Total Successes", results.TotalSuccesses},
		{"Total Slow Pings (> 1 sec)", results.TotalSlow},
		{"Total Errors", results.TotalErrors},
		{"Average Latency (µs)", results.AverageLatency.Microseconds()},
	})
	t.AppendSeparator()

	if len(results.SuccessUrls) > 0 {
		t.AppendRow(table.Row{"Successful URLs", ""})
		for _, url := range results.SuccessUrls {
			t.AppendRow(table.Row{"", url})
		}
		t.AppendSeparator()
	}

	if len(results.SlowUrls) > 0 {
		t.AppendRow(table.Row{"Slow URLs", ""})
		for _, url := range results.SlowUrls {
			t.AppendRow(table.Row{"", url})
		}
		t.AppendSeparator()
	}

	if len(results.FailedUrls) > 0 {
		t.AppendRow(table.Row{"Failed URLs", ""})
		for _, url := range results.FailedUrls {
			t.AppendRow(table.Row{"", url})
		}
	}

	t.Render()
}
