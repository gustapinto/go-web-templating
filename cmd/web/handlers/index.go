package handlers

import (
	"html/template"
	"net/http"
)

type Index struct {
	Views *template.Template
}

func (h Index) Root(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Messages []string
	}{
		Messages: []string{
			"Hello From Mars",
			"Or From The World?",
		},
	}
	err := h.Views.ExecuteTemplate(w, "index-view", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
