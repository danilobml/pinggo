package analizer

import (
	"testing"
	"time"

	"github.com/danilobml/pinggo/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_GenerateSummary_Success(t *testing.T) {
	input := models.PingerResponse{
		{
			Url: "www.test.com",
			StatusCode: 200,
			Latency: time.Duration(200 * time.Millisecond),
			Error: false,
		},
		{
			Url: "www.test2.com",
			StatusCode: 200,
			Latency: time.Duration(1000 * time.Millisecond),
			Error: false,
		},
		{
			Url: "www.fail-test.com",
			StatusCode: 400,
			Error: true,
		},
	}
	expected := models.SummaryResponse{
		TotalSuccesses: 2,
		TotalSlow: 1,
		TotalErrors: 1,
		AverageLatency: time.Duration(600 * time.Millisecond),
		SuccessUrls: []string{"www.test.com", "www.test2.com"},
		SlowUrls: []string{"www.test2.com"},
		FailedUrls: []string{"www.fail-test.com"},
	}

	actual := GenerateSummary(input)

	assert.Equal(t, expected, actual)
}
