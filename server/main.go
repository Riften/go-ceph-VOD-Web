package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template" //导入模版包
)



func myWeb(w http.ResponseWriter, r *http.Request) {

	t := template.New("index")

	t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>")

	data := map[string]string{
		"name":    "zeta",
		"someStr": "这是一个开始",
	}

	t.Execute(w, data)

	// fmt.Fprintln(w, "这是一个开始")
}

type testHandler struct {

}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.URL.Query() {
		fmt.Println("key:", k, ", value:", v[0])
	}

	for k, v := range r.PostForm {
		fmt.Println("key:", k, ", value:", v[0])
	}

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
		fmt.Fprintln(w, "这是一个开始")
	}
}

func main() {
	http.Handle("/", &testHandler{})
	fmt.Println("Start server and listen port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
