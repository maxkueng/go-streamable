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
	videoURL := "http://techslides.com/demos/samples/sample.wmv"
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
	assert.Equal(t, "streamable.com/ifjh", res.URL)
	assert.Contains(t, res.ThumbnailURL, "streamable.com/image/ifjh.jpg")
	assert.Equal(t, "", res.Message)

	expectedMp4 := VideoInfoFile{
		URL:    "streamable.com/video/mp4/ifjh.mp4",
		Width:  848,
		Height: 480,
	}
	assert.Contains(t, res.Files["mp4"].URL, expectedMp4.URL)

	expectedWebm := VideoInfoFile{
		URL:    "streamable.com/video/webm/ifjh.webm",
		Width:  848,
		Height: 480,
	}
	assert.Contains(t, res.Files["webm"].URL, expectedWebm.URL)
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
	assert.Equal(t, "streamable.com/ifjh", res.URL)
	assert.Contains(t, "streamable.com/image/ifjh.jpg", res.ThumbnailURL)
	assert.Equal(t, "", res.Message)
}

func Test_getVideoURL(t *testing.T) {
	u := getVideoURL("yolo")

	assert.Equal(t, videoURL+"/yolo", u)
}
