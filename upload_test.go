package streamable

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFilesDir = path.Join(".", "test-files")

func Test_Upload(t *testing.T) {
	var testFile = path.Join(testFilesDir, "cat-video.mp4")

	res, err := Upload(testFile)
	assert.Nil(t, err)
	assert.NotZero(t, res)
}
