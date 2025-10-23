package pinger

import (
	"testing"

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