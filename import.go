package streamable

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const importUrl string = apiUrl + "/import"

func ImportVideoFromUrl(videoUrl string) (VideoResponse, error) {
	return importVideoFromUrl(Credentials{}, videoUrl)
}

func ImportVideoFromUrlAuthenticated(creds Credentials, videoUrl string) (VideoResponse, error) {
	return importVideoFromUrl(creds, videoUrl)
}

func importVideoFromUrl(creds Credentials, videoUrl string) (VideoResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", getImportUrl(videoUrl), nil)
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
	if err != nil {
		return VideoResponse{}, err
	}

	body := bytesToString(bodyBytes)

	videoRes, err := videoResponseFromJson(body)
	if err != nil {
		return VideoResponse{}, err
	}

	return videoRes, nil

}

func getImportUrl(videoUrl string) string {
	parsedUrl, _ := url.Parse(importUrl)
	q := parsedUrl.Query()
	q.Set("url", videoUrl)
	parsedUrl.RawQuery = q.Encode()

	return parsedUrl.String()
}
