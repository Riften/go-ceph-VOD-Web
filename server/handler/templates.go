package handler

import (
	"fmt"
	"net/http"
	"strconv"
)


type VideoInfo struct {
	PosterUrl string
	VideoLength string
	VideoName string
	VideoUrl string
}
type VideoData struct {
	Videos []VideoInfo
}

type PlayData struct {
	VideoIndex string
	VideoPoster string
}

func (h *HttpHandler) rendIndex(w http.ResponseWriter) {
	fmt.Println("render index")
	tmpl, err := h.repo.FetchTemplate([]string{"index.html"})
	if err != nil {
		fmt.Println("Error when load template index.html: ", err)
		return
	}
	vs := h.repo.DataStore.Videos.List("")
	data := &VideoData{Videos: []VideoInfo{}}
	for _, v := range vs {
		info := VideoInfo{
			PosterUrl:   "/poster?name="+v.Poster,
			VideoLength: secondsToString(v.VideoLength),
			VideoName:   v.VideoName,
			VideoUrl:  fmt.Sprintf("http://localhost:8080/play.html?index=%d", v.Index),
		}
		data.Videos = append(data.Videos, info)
	}


	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		fmt.Println("Error when execute template index.html: ", err)
	}
}

func (h *HttpHandler) rendPlay(videoIndex int, w http.ResponseWriter) {
	fmt.Println( "rend play")
	tmpl, err := h.repo.FetchTemplate([]string{"play.html"})
	if err != nil {
		fmt.Println("Error when load template index.html: ", err)
		return
	}
	vs := h.repo.DataStore.Videos.List(fmt.Sprintf("ind=%d", videoIndex))
	if len(vs) > 0 {
		v:= vs[0]
		err = tmpl.ExecuteTemplate(w, "play", PlayData{
			VideoIndex:  strconv.Itoa(videoIndex),
			VideoPoster: v.Poster,
		})
	} else {
		fmt.Println("No video with index ", videoIndex)
	}
}