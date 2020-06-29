package handler

import (
	"fmt"
	"io"
	"main/db"
	"net/http"
	"os"
	"path"
)


type HttpHandler struct {
	//repoPath string
	repo *db.Repo
}

func NewHttpHandler(repo *db.Repo) *HttpHandler{
	return &HttpHandler{repo: repo}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error when parse request: %v\n", err)
		return
	}

	fmt.Println("Get request: "+r.Method)
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
	}
	//for k, v := range r.PostForm {
	//	//r.Body
	//	fmt.Println("key:", k, ", value:", v[0])
	//}

	fmt.Printf("Request: %s\n", r.RequestURI)
	if r.RequestURI == "/getVideos" {
		fmt.Println("Get video")
		file, err := os.Open("test1.mp4")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		BufferSize := 1024
		buffer := make([]byte, BufferSize)

		for {
			bytesread, err := file.Read(buffer)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Println("bytes read: ", bytesread)
			//fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
			w.Write(buffer)
		}
	}
}
