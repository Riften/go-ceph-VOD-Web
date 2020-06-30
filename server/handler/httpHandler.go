package handler

import (
	"bytes"
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"main/db"
	"net/http"
	"path"
	"strconv"
	"sync"
)


type HttpHandler struct {
	//repoPath string
	repo *db.Repo
	conn *rados.Conn
	cephPool string
	videoLock sync.Mutex
}

func NewHttpHandler(repo *db.Repo, conn *rados.Conn) *HttpHandler{
	return &HttpHandler{repo: repo, conn: conn, cephPool: "mytest"}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error when parse request: %v\n", err)
		return
	}

	fmt.Println("Get request: "+r.Method)
	//fmt.Printf("Request: %s\n", r.RequestURI)
	fmt.Println("\t",r.URL.Path)
	fmt.Println("query:")
	for k, v := range r.URL.Query() {
		fmt.Println("\t", "key:", k, ", value:", v[0])
	}
	values:= r.PostForm
	fmt.Println("values:")
	for k, v := range values {
		fmt.Println("\t", "key:", k, ", value:", v[0])
	}

	switch r.URL.Path {
	case "/favicon.ico":
		err = fetchFileToHttp(path.Join(h.repo.RepoPath(), "resource", "favicon.ico"), w)
		if err != nil {
			fmt.Println("Error when fetch resource/favicon.ico: "+err.Error())
		}
	case "/index.html":
		// rend main page
		h.rendIndex(w)
	case "/cephtest":
		// open a pool handle
		ioctx, err := h.conn.OpenIOContext("mytest")
		if err != nil {
			fmt.Println("Error when open ceph pool mytest: ", err)
			return
		}
		// write some data
		bytesIn := []byte("input data")

		fmt.Println("Try to write to ceph pool mytest.")
		err = ioctx.Write("obj", bytesIn, 0)

		// read the data back out
		bytesOut := make([]byte, len(bytesIn))
		_, err = ioctx.Read("obj", bytesOut, 0)
		if err != nil {
			fmt.Println("Error when read obj from ceph pool mytest: ", err)
			return
		}
		if !bytes.Equal(bytesIn, bytesOut) {
			fmt.Println("Output is not input!")
		}
		fmt.Println("Ceph works fine.")
	case "/addVideo":
		videoPath, ok1 := values["videoPath"]
		posterPath, ok2 := values["posterPath"]
		videoName, ok3 := values["videoName"]
		videoLength, ok4 := values["videoLength"]
		if !(ok1 && ok2 && ok3 && ok4) {
			fmt.Println("Error: missing value field.")
			return
		}

		videoLength64, err := strconv.ParseInt(videoLength[0], 10, 64)
		if err != nil {
			fmt.Println("Error: fail to transfer string ", videoLength[0], " to int64")
			return
		}
		err = h.addVideoCeph(videoPath[0], posterPath[0], videoName[0], videoLength64)
		if err != nil {
			fmt.Println("Error when add video: ", err)
			return
		}
		fmt.Println("Done add video.")
	case "/listVideo":
		fmt.Println("List all videos in db:")
		videos := h.repo.DataStore.Videos.List("")
		for _, v := range videos {
			fmt.Println("\tIndex: ", v.Index)
			fmt.Println("\t\tName: ",v.VideoName)
			fmt.Println("\t\tLength: ",v.VideoLength)
			fmt.Println("\t\tPoster: ", v.Poster)
		}
		fmt.Println("Done listVideo")
	case "/lastVideo":
		fmt.Println("Get the last video in db:")
		v := h.repo.DataStore.Videos.GetLast()
		if v!= nil {
			fmt.Println("\tIndex: ", v.Index)
			fmt.Println("\t\tName: ",v.VideoName)
			fmt.Println("\t\tLength: ",v.VideoLength)
			fmt.Println("\t\tPoster: ", v.Poster)
		} else {
			fmt.Println("Video db is empty.")
		}
		fmt.Println("Done lastVideo")
	case "/common.css":
		fmt.Println("Fetching common.css.")
		err = fetchFileToHttp(path.Join(h.repo.ResPath(), "css", "common.css"), w)
		if err != nil {
			fmt.Println("Error when fecth common.css: ", err)
		}
		fmt.Println("Done common.css")
	case "/poster":
		fmt.Println("Fetching poster")
		name, ok := values["name"]
		if !ok {
			fmt.Println("No field name in poster query.")
			return
		}
		err = fetchCephToHttp(h.conn, h.cephPool, name[0], w)
	default:
		fmt.Println("Unsupporter url path ", r.URL.Path)
	}

}
