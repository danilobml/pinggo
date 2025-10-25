package pinger

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_PingUrl_Success(t *testing.T) {
	url := "https://www.google.com"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second}
	result := pingUrl(ctx, client, url)

	assert.Equal(t, url, result.Url)
	assert.False(t, result.Error, "expected no error")
	assert.Greater(t, result.StatusCode, 0, "status code should be set")
	assert.IsType(t, time.Duration(0), result.Latency)
}

func Test_PingUrl_Error(t *testing.T) {
	url := "http://not-google.com"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second}
	result := pingUrl(ctx, client, url)

	assert.Equal(t, url, result.Url)
	assert.True(t, result.Error, "expected error")
	assert.Equal(t, 0, result.StatusCode)
	assert.IsType(t, time.Duration(0), result.Latency)
}
