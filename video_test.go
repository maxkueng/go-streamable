package streamable

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Video(t *testing.T) {
	res, err := Video("ifjh")

	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, 2, res.Status)
	assert.Equal(t, "ifjh", res.Shortcode)
	assert.Equal(t, "//cdn.streamable.com/video/mp4/ifjh", res.UrlRoot)
	assert.Equal(t, "streamable.com/ifjh", res.Url)
	assert.Equal(t, "//cdn.streamable.com/image/ifjh.jpg", res.ThumbnailUrl)
	assert.Equal(t, "", res.Message)

	expectedFormats := []string{"mp4", "webm"}
	assert.Equal(t, expectedFormats, res.Formats)

	expectedMp4 := VideoResponseFile{
		Url:    "//cdn.streamable.com/video/mp4/ifjh.mp4",
		Width:  848,
		Height: 480,
	}
	assert.Equal(t, expectedMp4, res.Files["mp4"])

	expectedWebm := VideoResponseFile{
		Url:    "//cdn.streamable.com/video/webm/ifjh.webm",
		Width:  848,
		Height: 480,
	}
	assert.Equal(t, expectedWebm, res.Files["webm"])
}

func Test_VideoAuthenticated_UsernamePassword(t *testing.T) {
	if os.Getenv("STREAMABLE_USERNAME") == "" || os.Getenv("STREAMABLE_PASSWORD") == "" {
		t.Skip("skipping test; $STREAMABLE_USERNAME or $STREAMABLE_PASSWORD not set")
	}

	creds := Credentials{
		Username: os.Getenv("STREAMABLE_USERNAME"),
		Password: os.Getenv("STREAMABLE_PASSWORD"),
	}

	res, err := VideoAuthenticated(creds, "ifjh")

	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, 2, res.Status)
	assert.Equal(t, "ifjh", res.Shortcode)
	assert.Equal(t, "//cdn.streamable.com/video/mp4/ifjh", res.UrlRoot)
	assert.Equal(t, "streamable.com/ifjh", res.Url)
	assert.Equal(t, "//cdn.streamable.com/image/ifjh.jpg", res.ThumbnailUrl)
	assert.Equal(t, "", res.Message)
}

func Test_getVideoUrl(t *testing.T) {
	u := getVideoUrl("yolo")

	assert.Equal(t, videoUrl+"/yolo", u)
}
