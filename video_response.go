package streamable

const (
	StatusUploading   = 0
	StatusProcessing  = 1
	StatusReady       = 2
	StatusUnavailable = 3
)

type VideoResponse struct {
	Status       int                          `json:"status"`
	Shortcode    string                       `json:"shortcode"`
	UrlRoot      string                       `json:"url_root"`
	Url          string                       `json:"url"`
	ThumbnailUrl string                       `json:"thumbnail_url"`
	Files        map[string]VideoResponseFile `json:"files"`
	Formats      []string                     `json:"formats"`
	Message      string                       `json:"message"`
}

type VideoResponseFile struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
