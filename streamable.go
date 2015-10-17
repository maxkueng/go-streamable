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
	apiURL    string = "https://api.streamable.com"
	importURL string = apiURL + "/import"
	uploadURL string = apiURL + "/upload"
	videoURL  string = apiURL + "/videos"
)

// UploadVideo uploads a video file located at filePath and returns a
// VideoInfo.
func UploadVideo(filePath string) (VideoInfo, error) {
	return uploadVideo(Credentials{}, filePath)
}

// UploadVideoAuthenticated uploads a video file located at filePath with
// authentication using the credentials provided in creds, and returns a
// VideoInfo.
func UploadVideoAuthenticated(creds Credentials, filePath string) (VideoInfo, error) {
	return uploadVideo(creds, filePath)
}

// UploadVideoFromURL uploads a video from a remote URL videoURL and returns a
// VideoInfo.
func UploadVideoFromURL(videoURL string) (VideoInfo, error) {
	return uploadVideoFromURL(Credentials{}, videoURL)
}

// UploadVideoFromURLAuthenticated uploads a video from a remote URL videoURL
// with authentication using the credentials provided in creds, and returns a
// VideoInfo.
func UploadVideoFromURLAuthenticated(creds Credentials, videoURL string) (VideoInfo, error) {
	return uploadVideoFromURL(creds, videoURL)
}

// GetVideo returns a VideoInfo with information about the video with the short
// code shortcode.
func GetVideo(shortcode string) (VideoInfo, error) {
	return getVideo(Credentials{}, shortcode)
}

// GetVideoAuthenticated returns a VideoInfo with information about the video
// with the short code shortcode with authentication using the credentials
// provides in creds.  This is useful to retrieve information about videos that
// aren't public.
func GetVideoAuthenticated(creds Credentials, shortcode string) (VideoInfo, error) {
	return getVideo(creds, shortcode)
}

func uploadVideoFromURL(creds Credentials, videoURL string) (VideoInfo, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", getImportURL(videoURL), nil)
	if err != nil {
		return VideoInfo{}, err
	}

	if creds.Username != "" && creds.Password != "" {
		req.SetBasicAuth(creds.Username, creds.Password)
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return VideoInfo{}, err
	}

	if res.StatusCode != http.StatusOK {
		return VideoInfo{}, fmt.Errorf("not found")
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return VideoInfo{}, err
	}

	body := bytesToString(bodyBytes)

	videoRes, err := videoResponseFromJSON(body)
	if err != nil {
		return VideoInfo{}, err
	}

	return videoRes, nil

}

func uploadVideo(creds Credentials, filePath string) (VideoInfo, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return VideoInfo{}, err
	}

	var buf bytes.Buffer

	multipartWriter := multipart.NewWriter(&buf)

	fileHandle, err := os.Open(filePath)
	if err != nil {
		return VideoInfo{}, err
	}

	fileWriter, err := multipartWriter.CreateFormFile("file", filePath)
	if err != nil {
		return VideoInfo{}, err
	}

	_, err = io.Copy(fileWriter, fileHandle)
	if err != nil {
		return VideoInfo{}, err
	}

	multipartWriter.Close()

	req, err := http.NewRequest("POST", uploadURL, &buf)
	if err != nil {
		return VideoInfo{}, err
	}

	if creds.Username != "" && creds.Password != "" {
		req.SetBasicAuth(creds.Username, creds.Password)
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return VideoInfo{}, err
	}

	if res.StatusCode != http.StatusOK {
		return VideoInfo{}, fmt.Errorf("upload failed")
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return VideoInfo{}, err
	}

	body := bytesToString(bodyBytes)

	videoRes, err := videoResponseFromJSON(body)
	if err != nil {
		return VideoInfo{}, err
	}

	return videoRes, nil
}

func getVideo(creds Credentials, shortcode string) (VideoInfo, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", getVideoURL(shortcode), nil)
	if err != nil {
		return VideoInfo{}, err
	}

	if creds.Username != "" && creds.Password != "" {
		req.SetBasicAuth(creds.Username, creds.Password)
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return VideoInfo{}, err
	}

	if res.StatusCode != http.StatusOK {
		return VideoInfo{}, fmt.Errorf("not found")
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return VideoInfo{}, err
	}

	body := bytesToString(bodyBytes)

	videoRes, err := videoResponseFromJSON(body)
	if err != nil {
		return VideoInfo{}, err
	}

	videoRes.Shortcode = shortcode

	return videoRes, nil
}

func getImportURL(videoURL string) string {
	parsedURL, _ := url.Parse(importURL)
	q := parsedURL.Query()
	q.Set("url", videoURL)
	parsedURL.RawQuery = q.Encode()

	return parsedURL.String()
}

func getVideoURL(shortcode string) string {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return ""
	}

	parsedURL.Path += "/" + shortcode

	return parsedURL.String()
}
