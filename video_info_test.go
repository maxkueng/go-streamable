package streamable

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VideoInfo(t *testing.T) {
	jsonStr := `{
			"status": 2,
			"files": {
				"mp4": {
					"url": "//cdn.streamable.com/video/mp4/ifjh.mp4",
					"width": 848,
					"height": 480
				},
				"webm": {
					"url": "//cdn.streamable.com/video/webm/ifjh.webm",
					"width": 848,
					"height": 480
				}
			},
			"url": "streamable.com/ifjh",
			"thumbnail_url": "//cdn.streamable.com/image/ifjh.jpg",
			"message": null
		}`

	res := VideoInfo{}
	err := json.Unmarshal([]byte(jsonStr), &res)

	assert.Nil(t, err)

	assert.Equal(t, 2, res.Status)
	assert.Equal(t, "", res.Shortcode)
	assert.Equal(t, "streamable.com/ifjh", res.URL)
	assert.Equal(t, "//cdn.streamable.com/image/ifjh.jpg", res.ThumbnailURL)
	assert.Equal(t, "", res.Message)

	expectedMp4 := VideoInfoFile{
		URL:    "streamable.com/video/mp4/ifjh.mp4",
		Width:  848,
		Height: 480,
	}
	resMp4 := res.Files["mp4"]
	assert.Contains(t, resMp4.URL, expectedMp4.URL)
	assert.Equal(t, resMp4.Height, expectedMp4.Height)
	assert.Equal(t, resMp4.Width, expectedMp4.Width)

	expectedWebm := VideoInfoFile{
		URL:    "//cdn.streamable.com/video/webm/ifjh.webm",
		Width:  848,
		Height: 480,
	}
	assert.Equal(t, expectedWebm, res.Files["webm"])
}

func Test_videoResponseFromJSON(t *testing.T) {
	jsonStr := `{
			"status": 2,
			"files": {
				"mp4": {
					"url": "//cdn.streamable.com/video/mp4/ifjh.mp4",
					"width": 848,
					"height": 480
				},
				"webm": {
					"url": "//cdn.streamable.com/video/webm/ifjh.webm",
					"width": 848,
					"height": 480
				}
			},
			"url_root": "//cdn.streamable.com/video/mp4/ifjh",
			"url": "streamable.com/ifjh",
			"thumbnail_url": "//cdn.streamable.com/image/ifjh.jpg",
			"formats": [
				"mp4",
				"webm"
			],
			"message": null
		}`

	res, err := videoResponseFromJSON(jsonStr)

	assert.Nil(t, err)

	assert.Equal(t, 2, res.Status)
	assert.Equal(t, "", res.Shortcode)
	assert.Equal(t, "streamable.com/ifjh", res.URL)
	assert.Contains(t, "//cdn.streamable.com/image/ifjh.jpg", res.ThumbnailURL)
	assert.Equal(t, "", res.Message)

	expectedMp4 := VideoInfoFile{
		URL:    "//cdn.streamable.com/video/mp4/ifjh.mp4",
		Width:  848,
		Height: 480,
	}
	resMp4 := res.Files["mp4"]
	assert.Contains(t, resMp4.URL, expectedMp4.URL)
	assert.Equal(t, expectedMp4.Height, resMp4.Height)
	assert.Equal(t, expectedMp4.Width, resMp4.Width)

	expectedWebm := VideoInfoFile{
		URL:    "//cdn.streamable.com/video/webm/ifjh.webm",
		Width:  848,
		Height: 480,
	}
	assert.Equal(t, expectedWebm, res.Files["webm"])
}
