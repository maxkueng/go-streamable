package streamable

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const videoUrl string = apiUrl + "/videos"

func Video(shortcode string) (VideoResponse, error) {
	return video(Credentials{}, shortcode)
}

func VideoAuthenticated(creds Credentials, shortcode string) (VideoResponse, error) {
	return video(creds, shortcode)
}

func video(creds Credentials, shortcode string) (VideoResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", getVideoUrl(shortcode), nil)
	if err != nil {
		return VideoResponse{}, err
	}

	if creds.Username != "" && creds.Password != "" {
		req.SetBasicAuth(creds.Username, creds.Password)
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return VideoResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return VideoResponse{}, fmt.Errorf("not found")
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	body := bytesToString(bodyBytes)

	videoRes, err := videoResponseFromJson(body)
	if err != nil {
		return VideoResponse{}, err
	}

	videoRes.Shortcode = shortcode

	return videoRes, nil
}

func getVideoUrl(shortcode string) string {
	parsedUrl, err := url.Parse(videoUrl)
	if err != nil {
		return ""
	}

	parsedUrl.Path += "/" + shortcode

	return parsedUrl.String()
}
