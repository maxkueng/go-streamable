package streamable

import "encoding/json"

const (
	StatusUploading   = 0
	StatusProcessing  = 1
	StatusReady       = 2
	StatusUnavailable = 3
)

type VideoResponse struct {
	Status       int                          `json:"status"`
	Shortcode    string                       `json:"shortcode"`
	URLRoot      string                       `json:"url_root"`
	URL          string                       `json:"url"`
	ThumbnailURL string                       `json:"thumbnail_url"`
	Files        map[string]VideoResponseFile `json:"files"`
	Formats      []string                     `json:"formats"`
	Message      string                       `json:"message"`
}

type VideoResponseFile struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func videoResponseFromJSON(jsonStr string) (VideoResponse, error) {
	res := VideoResponse{}
	parseErr := json.Unmarshal([]byte(jsonStr), &res)
	if parseErr != nil {
		return VideoResponse{}, parseErr
	}

	return res, nil
}
