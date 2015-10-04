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

func Test_Upload_UsernamePassword(t *testing.T) {
	var testFile = path.Join(testFilesDir, "cat-video.mp4")

	creds := Credentials{
		Username: "gostreamabletest",
		Password: "gostreamabletest0=",
	}

	res, err := UploadAuthenticated(creds, testFile)
	assert.Nil(t, err)
	assert.NotZero(t, res)
}
