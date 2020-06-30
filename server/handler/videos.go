package handler

import (
	"errors"
	"fmt"
	"main/db"
)

/*
type Video struct {
	Index int `json:"index"`
	VideoName string `json:"videoName"`
	BlockNum int `json:"blockNum"`
	Poster string `json:"poster"`
	VideoLength int64 `json:"videoLength"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
}
 */
// Add a new video to ceph.
// The whole video would be treated as a single ceph object.
// The block id of this video would be video_Index
// The block id of the poster would be poster_Index
func (h *HttpHandler) addVideoCeph(videoPath string, posterPath string, videoName string, videoLength int64) error {
	h.videoLock.Lock()
	defer h.videoLock.Unlock()

	var videoInd int
	lastVideo := h.repo.DataStore.Videos.GetLast()
	if lastVideo == nil {
		fmt.Println("Adding the first video")
		videoInd = 0
	} else {
		fmt.Println("Adding video with index ", lastVideo.Index + 1)
		videoInd = lastVideo.Index + 1
	}

	if !fileExists(videoPath) {
		fmt.Println(videoPath, " not exists.")
		return errors.New(videoPath + " not exists")
	}
	if !fileExists(posterPath) {
		fmt.Println(posterPath, " not exists")
		return errors.New(posterPath + " not exists")
	}
	videoObject := fmt.Sprintf("video_%d", videoInd)
	posterObject := fmt.Sprintf("poster_%d", videoInd)
	err := fetchFileToCeph(videoPath, h.conn, h.cephPool, videoObject)
	if err != nil {
		fmt.Println("Error occur when fetch ", videoPath, "to ceph object ", videoObject, ": ", err)
		return err
	}
	err = fetchFileToCeph(posterPath, h.conn, h.cephPool, posterObject)
	if err != nil {
		fmt.Println("Error occur when fetch ", posterPath, " to ceph object ", posterObject, ": ", err)
		return err
	}

	v := &db.Video{
		Index:       videoInd,
		VideoName:   videoName,
		BlockNum:    1,
		Poster:      posterObject,
		VideoLength: videoLength,
		Created:     0,
		Updated:     0,
	}
	err = h.repo.DataStore.Videos.Add(v)
	if err != nil {
		fmt.Println("Error occur when add video ", videoName, " to datastore: ", err)
		return err
	}
	return nil
}