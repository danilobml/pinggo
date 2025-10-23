package parser

import (
	"testing"

	"github.com/danilobml/pinggo/internal/errs"
	"github.com/stretchr/testify/assert"
)

func Test_GetUrlsFromFile_Success(t *testing.T) {
	filePath := "./test.txt"
	expected := []string{"http://www.test.com", "http://www.test2.com"}
	
	actual, err := GetUrlsFromFile(filePath)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func Test_GetUrlsFromFile_Error_NotTxt(t *testing.T) {
	filePath := "./test.pdf"
	_, err := GetUrlsFromFile(filePath)
	assert.ErrorIs(t, err, errs.ErrInvalidInputFile)
}
