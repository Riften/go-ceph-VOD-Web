// Client related command
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func sendRequest() error{

	//testReq, err := http.NewRequest("POST", "/cmd", nil)
	apiUrl := localhost
	resource := "test"
	data := url.Values{}
	data.Set("name", "xiaohua")
	data.Set("id", "654321")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource

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
	fmt.Println("[client.Do] request2 sent successfully.")
	return nil
}