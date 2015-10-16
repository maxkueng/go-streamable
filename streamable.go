// Package streamable provides functions to interact with the
// https://streamable.com/ API.
package streamable

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

const (
	apiUrl    string = "https://api.streamable.com"
	importUrl string = apiUrl + "/import"
	uploadUrl string = apiUrl + "/upload"
	videoUrl  string = apiUrl + "/videos"
)

// UploadVideo uploads a video file located at filePath and returns a
// VideoResponse.
func UploadVideo(filePath string) (VideoResponse, error) {
	return uploadVideo(Credentials{}, filePath)
}

// UploadVideoAuthenticated uploads a video file located at filePath with
// authentication using the credentials provided in creds, and returns a
// VideoResponse.
func UploadVideoAuthenticated(creds Credentials, filePath string) (VideoResponse, error) {
	return uploadVideo(creds, filePath)
}

// ImportVideoFromUrl uploads a video from a remote URL videoUrl and returns a
// VideoResponse.
func ImportVideoFromUrl(videoUrl string) (VideoResponse, error) {
	return importVideoFromUrl(Credentials{}, videoUrl)
}

// ImportVideoFromUrlAuthenticated uploads a video from a remote URL videoUrl
// with authentication using the credentials provided in creds, and returns a
// VideoResponse.
func ImportVideoFromUrlAuthenticated(creds Credentials, videoUrl string) (VideoResponse, error) {
	return importVideoFromUrl(creds, videoUrl)
}

// GetVideo returns a VideoResponse with information about the video with the
// short code shortcode.
func GetVideo(shortcode string) (VideoResponse, error) {
	return getVideo(Credentials{}, shortcode)
}

// GetVideo returns a VideoResponse with information about the video with the
// short code shortcode with authentication using the credentials provides in
// creds.
// This is useful to retrieve information about videos that aren't public.
func GetVideoAuthenticated(creds Credentials, shortcode string) (VideoResponse, error) {
	return getVideo(creds, shortcode)
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

func uploadVideo(creds Credentials, filePath string) (VideoResponse, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return VideoResponse{}, err
	}

	var buf bytes.Buffer

	multipartWriter := multipart.NewWriter(&buf)

	fileHandle, err := os.Open(filePath)
	if err != nil {
		return VideoResponse{}, err
	}

	fileWriter, err := multipartWriter.CreateFormFile("file", filePath)
	if err != nil {
		return VideoResponse{}, err
	}

	_, err = io.Copy(fileWriter, fileHandle)
	if err != nil {
		return VideoResponse{}, err
	}

	multipartWriter.Close()

	req, err := http.NewRequest("POST", uploadUrl, &buf)
	if err != nil {
		return VideoResponse{}, err
	}

	if creds.Username != "" && creds.Password != "" {
		req.SetBasicAuth(creds.Username, creds.Password)
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return VideoResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return VideoResponse{}, fmt.Errorf("upload failed")
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

func getVideo(creds Credentials, shortcode string) (VideoResponse, error) {
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
	if err != nil {
		return VideoResponse{}, err
	}

	body := bytesToString(bodyBytes)

	videoRes, err := videoResponseFromJson(body)
	if err != nil {
		return VideoResponse{}, err
	}

	videoRes.Shortcode = shortcode

	return videoRes, nil
}

func getImportUrl(videoUrl string) string {
	parsedUrl, _ := url.Parse(importUrl)
	q := parsedUrl.Query()
	q.Set("url", videoUrl)
	parsedUrl.RawQuery = q.Encode()

	return parsedUrl.String()
}

func getVideoUrl(shortcode string) string {
	parsedUrl, err := url.Parse(videoUrl)
	if err != nil {
		return ""
	}

	parsedUrl.Path += "/" + shortcode

	return parsedUrl.String()
}
