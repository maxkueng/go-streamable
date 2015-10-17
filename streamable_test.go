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

	c := New()
	res, err := c.UploadVideo(testFile)

	assert.Nil(t, err)
	assert.NotZero(t, res)
}

func Test_UploadVideo_NonExistentFile(t *testing.T) {
	testFile := path.Join(testFilesDir, "not-exists.mp4")

	c := New()
	res, err := c.UploadVideo(testFile)

	assert.NotNil(t, err)
	assert.Zero(t, res)
}

func Test_UploadVideo_Authenticated_UsernamePassword(t *testing.T) {
	if os.Getenv("STREAMABLE_USERNAME") == "" || os.Getenv("STREAMABLE_PASSWORD") == "" {
		t.Skip("skipping test; $STREAMABLE_USERNAME or $STREAMABLE_PASSWORD not set")
	}

	testFile := path.Join(testFilesDir, "cat-video.mp4")

	c := New()
	c.SetCredentials(os.Getenv("STREAMABLE_USERNAME"), os.Getenv("STREAMABLE_PASSWORD"))
	res, err := c.UploadVideo(testFile)

	assert.Nil(t, err)
	assert.NotZero(t, res)
}

func Test_UploadVideoFromURL(t *testing.T) {
	videoURL := "https://archive.org/download/Windows7WildlifeSampleVideo/Wildlife.wmv"
	c := New()
	res, err := c.UploadVideoFromURL(videoURL)

	assert.Nil(t, err)
	assert.Equal(t, 1, res.Status)
	assert.NotEqual(t, "", res.Shortcode)
}

func Test_UploadVideoFromURL_Authenticated_UsernamePassword(t *testing.T) {
	if os.Getenv("STREAMABLE_USERNAME") == "" || os.Getenv("STREAMABLE_PASSWORD") == "" {
		t.Skip("skipping test; $STREAMABLE_USERNAME or $STREAMABLE_PASSWORD not set")
	}

	c := New()
	c.SetCredentials(os.Getenv("STREAMABLE_USERNAME"), os.Getenv("STREAMABLE_PASSWORD"))

	videoURL := "https://archive.org/download/Windows7WildlifeSampleVideo/Wildlife.wmv"
	res, err := c.UploadVideoFromURL(videoURL)

	assert.Nil(t, err)
	assert.Equal(t, 1, res.Status)
	assert.NotEqual(t, "", res.Shortcode)
}

func Test_GetVideo(t *testing.T) {
	c := New()
	res, err := c.GetVideo("ifjh")

	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, 2, res.Status)
	assert.Equal(t, "ifjh", res.Shortcode)
	assert.Equal(t, "//cdn.streamable.com/video/mp4/ifjh", res.URLRoot)
	assert.Equal(t, "streamable.com/ifjh", res.URL)
	assert.Equal(t, "//cdn.streamable.com/image/ifjh.jpg", res.ThumbnailURL)
	assert.Equal(t, "", res.Message)

	expectedFormats := []string{"mp4", "webm"}
	assert.Equal(t, expectedFormats, res.Formats)

	expectedMp4 := VideoInfoFile{
		URL:    "//cdn.streamable.com/video/mp4/ifjh.mp4",
		Width:  848,
		Height: 480,
	}
	assert.Equal(t, expectedMp4, res.Files["mp4"])

	expectedWebm := VideoInfoFile{
		URL:    "//cdn.streamable.com/video/webm/ifjh.webm",
		Width:  848,
		Height: 480,
	}
	assert.Equal(t, expectedWebm, res.Files["webm"])
}

func Test_GetVideo_Authenticated_UsernamePassword(t *testing.T) {
	if os.Getenv("STREAMABLE_USERNAME") == "" || os.Getenv("STREAMABLE_PASSWORD") == "" {
		t.Skip("skipping test; $STREAMABLE_USERNAME or $STREAMABLE_PASSWORD not set")
	}

	c := New()
	c.SetCredentials(os.Getenv("STREAMABLE_USERNAME"), os.Getenv("STREAMABLE_PASSWORD"))
	res, err := c.GetVideo("ifjh")

	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, 2, res.Status)
	assert.Equal(t, "ifjh", res.Shortcode)
	assert.Equal(t, "//cdn.streamable.com/video/mp4/ifjh", res.URLRoot)
	assert.Equal(t, "streamable.com/ifjh", res.URL)
	assert.Equal(t, "//cdn.streamable.com/image/ifjh.jpg", res.ThumbnailURL)
	assert.Equal(t, "", res.Message)
}

func Test_getVideoURL(t *testing.T) {
	u := getVideoURL("yolo")

	assert.Equal(t, videoURL+"/yolo", u)
}
