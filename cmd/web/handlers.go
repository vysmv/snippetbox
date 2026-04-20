package main

import (
	"fmt"
	"html/template" // New import
	"strconv"
	// New import
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler is now a method against the application
		// struct it can access its fields, including the structured logger. We'll
		// use this to create a log entry at Error level containing the error
		// message, also including the request method and URI as attributes to
		// assist with debugging.
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// And we also need to update the code here to use the structured logger
		// too.
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) listSnippets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly list snippets"))
}

func (app *application) newSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly form for create snippet"))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create snippet by ID"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display snippet by ID"))
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) editSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for edit snippet"))
}

func (app *application) updateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update snippet by ID"))
}

func (app *application) deleteSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete snippet by ID"))
}
