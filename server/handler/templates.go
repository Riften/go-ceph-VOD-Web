package handler

import (
	"fmt"
	"net/http"
)

func (h *HttpHandler) rendIndex(w http.ResponseWriter) {
	tmpl, err := h.repo.FetchTemplate([]string{"layout.html"})
	if err != nil {
		fmt.Println("Error when load template layout.html: ", err)
		return
	}
	err = tmpl.Execute(w, "模板测试")
	if err != nil {
		fmt.Println("Error when execute template layout.html: ", err)
	}
}