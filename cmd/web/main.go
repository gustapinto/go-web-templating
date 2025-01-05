package main

import (
	"embed"
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/gustapinto/go-web-templating/cmd/web/handlers"
)

//go:embed all:views
var viewsFS embed.FS

func main() {
	patterns := []string{
		"views/*.tmpl",
		"views/partials/*.tmpl",
	}
	views, err := template.New("").ParseFS(viewsFS, patterns...)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	index := handlers.Index{Views: views}
	mux.HandleFunc("GET /", index.Messages)
	mux.HandleFunc("POST /", index.MessagesForm)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server listening at address %s\n", listener.Addr().String())

	if err := http.Serve(listener, mux); err != nil {
		panic(err)
	}
}
