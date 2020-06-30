package handler

import (
	"fmt"
	"net/http"
)

func (h *HttpHandler) rendIndex(w http.ResponseWriter) {
	fmt.Println("render index")
	tmpl, err := h.repo.FetchTemplate([]string{"layout.html"})
	if err != nil {
		fmt.Println("Error when load template layout.html: ", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", "Hello world")
	if err != nil {
		fmt.Println("Error when execute template layout.html: ", err)
	}
}