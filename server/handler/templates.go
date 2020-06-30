package handler

import (
	"fmt"
	"net/http"
)


type VideoInfo struct {
	PosterUrl string
	VideoLength string
}
type VideoData struct {
	Videos []VideoInfo
}

func (h *HttpHandler) rendIndex(w http.ResponseWriter) {
	fmt.Println("render index")
	tmpl, err := h.repo.FetchTemplate([]string{"index.html"})
	if err != nil {
		fmt.Println("Error when load template index.html: ", err)
		return
	}

	data := VideoData{Videos: []VideoInfo{
		{
			PosterUrl:   "/poster?name=poster_0",
			VideoLength: secondsToString(1000),
		},
		{
			PosterUrl:   "/poster?name=poster_1",
			VideoLength: secondsToString(2000),
		},
	}}

	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		fmt.Println("Error when execute template index.html: ", err)
	}
}