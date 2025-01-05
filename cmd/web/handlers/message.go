package handlers

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"

	"github.com/gustapinto/go-web-templating/cmd/web/dto/request"
	"github.com/gustapinto/go-web-templating/internal/message"
)

type Message struct {
	Views   *template.Template
	Service message.Service
}

func NewMessage(views *template.Template, service message.Service) Message {
	return Message{
		Views:   views,
		Service: service,
	}
}

type contextKey string

var errorskey contextKey = "errors"

func (h *Message) MessagesListView(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Messages []string
		Errors   []string
	}{
		Messages: h.Service.GetAll(),
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

func (h *Message) MessagesForm(w http.ResponseWriter, r *http.Request) {
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
		h.MessagesListView(w, r)
		return
	}

	if err := h.Service.Create(req.Message); err != nil {
		ctx := context.WithValue(r.Context(), errorskey, []string{err.Error()})
		r = r.WithContext(ctx)
		h.MessagesListView(w, r)
		return
	}

	h.MessagesListView(w, r)
}
