package streamable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const uploadUrl string = apiUrl + "/upload"

func Upload(filePath string) (VideoResponse, error) {
	return upload(Credentials{}, filePath)
}

func UploadAuthenticated(creds Credentials, filePath string) (VideoResponse, error) {
	return upload(creds, filePath)
}

func upload(creds Credentials, filePath string) (VideoResponse, error) {
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

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return VideoResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return VideoResponse{}, fmt.Errorf("upload failed")
	}

	jsonRes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return VideoResponse{}, err
	}

	videoRes := VideoResponse{}
	err = json.Unmarshal(jsonRes, &videoRes)
	if err != nil {
		return VideoResponse{}, err
	}

	return videoRes, nil
}
