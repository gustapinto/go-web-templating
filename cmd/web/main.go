package main

import (
	"embed"
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/gustapinto/go-web-templating/cmd/web/handlers"
	"github.com/gustapinto/go-web-templating/internal/message"
	message_repository "github.com/gustapinto/go-web-templating/internal/message/repository"
)

//go:embed all:views
var viewsFS embed.FS

//go:embed all:assets
var assetsFS embed.FS

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
	mux.Handle("GET /assets/", http.FileServerFS(assetsFS))

	messageService := message.NewMessage(message_repository.NewInMemory())
	messageHandler := handlers.NewMessage(views, messageService)

	mux.HandleFunc("GET /{$}", messageHandler.MessagesListView)
	mux.HandleFunc("POST /{$}", messageHandler.MessagesForm)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server listening at address %s\n", listener.Addr().String())

	if err := http.Serve(listener, mux); err != nil {
		panic(err)
	}
}
