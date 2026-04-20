package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	 app := &application{
        logger: logger,
    }

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// List entities
	mux.HandleFunc("GET /snippets", app.listSnippets)
	// Form of creation
	mux.HandleFunc("GET /snippets/create", app.newSnippetForm)
	// Create one entity
	mux.HandleFunc("POST /snippets", app.createSnippet)
	// View one entity
	mux.HandleFunc("GET /snippets/{id}", app.showSnippet)
	// Form of edit
	mux.HandleFunc("GET /snippets/{id}/edit", app.editSnippetForm)
	// Update one entity
	mux.HandleFunc("PATCH /snippets/{id}", app.updateSnippet)
	// Delete one entity
	mux.HandleFunc("DELETE /snippets/{id}", app.deleteSnippet)

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
