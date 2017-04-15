package streamable

import "encoding/json"

// Represents a video resource's processing status.
const (
	StatusUploading   = 0
	StatusProcessing  = 1
	StatusReady       = 2
	StatusUnavailable = 3
)

// A VideoInfo represents a video resource.
type VideoInfo struct {
	Status       int                      `json:"status"`
	Shortcode    string                   `json:"shortcode"`
	URL          string                   `json:"url"`
	ThumbnailURL string                   `json:"thumbnail_url"`
	Files        map[string]VideoInfoFile `json:"files"`
	Message      string                   `json:"message"`
}

// A VideoInfoFile represents a single file of a VideoInfo.
type VideoInfoFile struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func videoResponseFromJSON(jsonStr string) (VideoInfo, error) {
	res := VideoInfo{}
	parseErr := json.Unmarshal([]byte(jsonStr), &res)
	if parseErr != nil {
		return VideoInfo{}, parseErr
	}

	return res, nil
}
