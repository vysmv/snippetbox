package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)

	// List entities
	mux.HandleFunc("GET /snippets", listSnippets)
	// Form of creation
	mux.HandleFunc("GET /snippets/create", newSnippetForm)
	// Create one entity
	mux.HandleFunc("POST /snippets", createSnippet)
	// View one entity
	mux.HandleFunc("GET /snippets/{id}", showSnippet)
	// Form of edit
	mux.HandleFunc("GET /snippets/{id}/edit", editSnippetForm)
	// Update one entity 
	mux.HandleFunc("PATCH /snippets/{id}", updateSnippet)
	// Delete one entity
	mux.HandleFunc("DELETE /snippets/{id}", deleteSnippet)

	log.Print("staring server on:4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
