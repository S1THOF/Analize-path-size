package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize_file1(t *testing.T) {
	res, err := GetPathSize("testdata/file1.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "62289B	testdata/file1.csv")
}

func TestGetSize_file2(t *testing.T) {
	res, err := GetPathSize("testdata/file2.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "62388B	testdata/file2.csv")
}

func TestGetSize_testdir(t *testing.T) {
	res, err := GetPathSize("testdata/testdir", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "124677B	testdata/testdir")
}
