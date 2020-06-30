package handler

import (
	"fmt"
	"net/http"
)

func (h *HttpHandler) rendIndex(w http.ResponseWriter) {
	fmt.Println("render index")
	tmpl, err := h.repo.FetchTemplate([]string{"index.html"})
	if err != nil {
		fmt.Println("Error when load template index.html: ", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index", "Hello world")
	if err != nil {
		fmt.Println("Error when execute template index.html: ", err)
	}
}