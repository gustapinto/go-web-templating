package handlers

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"

	"github.com/gustapinto/go-web-templating/cmd/web/dto/request"
)

type Index struct {
	Views    *template.Template
	messages []string
}

type contextKey string

var errorskey contextKey = "errors"

func (h *Index) Messages(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Messages []string
		Errors   []string
	}{
		Messages: h.messages,
		Errors:   nil,
	}

	if ctxErrs := r.Context().Value(errorskey); ctxErrs != nil {
		fmt.Println(ctxErrs)
		if errs, ok := ctxErrs.([]string); ok && len(errs) > 0 {
			data.Errors = errs
		}
	}

	err := h.Views.ExecuteTemplate(w, "index-view", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}

func (h *Index) MessagesForm(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	req := request.NewMessageForm(values)

	if err := req.Validate(); err != nil {
		ctx := context.WithValue(r.Context(), errorskey, []string{err.Error()})
		r = r.WithContext(ctx)
		h.Messages(w, r)
		return
	}

	h.messages = append(h.messages, req.Message)

	h.Messages(w, r)
}
