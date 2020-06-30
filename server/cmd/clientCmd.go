// Client related command
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func touchCeph() error {
	return sendRequest("cephtest", nil)
}

func addVideo(videoPath string, posterPath string, videoName string, videoLength int) error {
	req := "addVideo"
	values := map[string]string {
		"videoPath": videoPath,
		"posterPath" : posterPath,
		"videoName": videoName,
		"videoLength": fmt.Sprintf("%d", videoLength),
	}
	return sendRequest(req, values)
}

func sendRequest(path string, values map[string]string) error{
	//testReq, err := http.NewRequest("POST", "/cmd", nil)
	apiUrl := localhost
	//resource := "test"
	data := url.Values{}
	if values != nil {
		for k, v := range values {
			data.Set(k, v)
		}
	}

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = path

	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	return nil
}