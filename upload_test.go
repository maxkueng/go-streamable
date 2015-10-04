package streamable

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFilesDir = path.Join(".", "test-files")

func Test_UploadVideo(t *testing.T) {
	testFile := path.Join(testFilesDir, "cat-video.mp4")

	res, err := UploadVideo(testFile)
	assert.Nil(t, err)
	assert.NotZero(t, res)
}

func Test_UploadVideo_NonExistentFile(t *testing.T) {
	testFile := path.Join(testFilesDir, "not-exists.mp4")

	res, err := UploadVideo(testFile)
	assert.NotNil(t, err)
	assert.Zero(t, res)
}

func Test_UploadVideoAuthenticated_UsernamePassword(t *testing.T) {
	if os.Getenv("STREAMABLE_USERNAME") == "" || os.Getenv("STREAMABLE_PASSWORD") == "" {
		t.Skip("skipping test; $STREAMABLE_USERNAME or $STREAMABLE_PASSWORD not set")
	}

	creds := Credentials{
		Username: os.Getenv("STREAMABLE_USERNAME"),
		Password: os.Getenv("STREAMABLE_PASSWORD"),
	}

	testFile := path.Join(testFilesDir, "cat-video.mp4")

	res, err := UploadVideoAuthenticated(creds, testFile)
	assert.Nil(t, err)
	assert.NotZero(t, res)
}
