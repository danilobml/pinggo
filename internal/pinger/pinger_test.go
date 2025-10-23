package pinger

import (
	"testing"
	"time"

	"github.com/danilobml/pinggo/internal/errs"
	"github.com/stretchr/testify/assert"
)

func Test_GetUrlsFromFile_Success(t *testing.T) {
	filePath := "./test.txt"
	expected := []string{"test.com", "test2.com"}
	
	actual, err := getUrlsFromFile(filePath)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func Test_GetUrlsFromFile_Error_NotTxt(t *testing.T) {
	filePath := "./test.pdf"
	_, err := getUrlsFromFile(filePath)
	assert.ErrorIs(t, err, errs.ErrInvalidInputFile)
}

func Test_PingUrl_Success(t *testing.T) {
	url := "http://www.google.com"
	expected := PingUrlResponse{
		StatusCode: 200,
		Error: false,
	}
	
	actual, err := pingUrl(url)
	assert.NoError(t, err)

	assert.Equal(t, expected.StatusCode, actual.StatusCode)
	assert.Equal(t, expected.Error, actual.Error)
	assert.IsType(t, time.Duration(0), actual.Latency)
}

func Test_PingUrl_Error(t *testing.T) {
	url := "http://www.google.coms"
	expected := PingUrlResponse{
		Error: true,
	}
	
	actual, err := pingUrl(url)
	assert.ErrorIs(t, err, errs.ErrPingFailed)

	assert.Equal(t, expected, actual)
}
