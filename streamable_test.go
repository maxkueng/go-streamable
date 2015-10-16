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

func Test_ImportVideo(t *testing.T) {
	videoUrl := "https://archive.org/download/Windows7WildlifeSampleVideo/Wildlife.wmv"
	res, err := ImportVideoFromUrl(videoUrl)

	assert.Nil(t, err)
	assert.Equal(t, 1, res.Status)
	assert.NotEqual(t, "", res.Shortcode)
}

func Test_ImportVideoFromUrlAuthenticated_UsernamePassword(t *testing.T) {
	if os.Getenv("STREAMABLE_USERNAME") == "" || os.Getenv("STREAMABLE_PASSWORD") == "" {
		t.Skip("skipping test; $STREAMABLE_USERNAME or $STREAMABLE_PASSWORD not set")
	}

	creds := Credentials{
		Username: os.Getenv("STREAMABLE_USERNAME"),
		Password: os.Getenv("STREAMABLE_PASSWORD"),
	}

	videoUrl := "https://archive.org/download/Windows7WildlifeSampleVideo/Wildlife.wmv"
	res, err := ImportVideoFromUrlAuthenticated(creds, videoUrl)

	assert.Nil(t, err)
	assert.Equal(t, 1, res.Status)
	assert.NotEqual(t, "", res.Shortcode)
}
