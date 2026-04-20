package main

import (
	"html/template" // New import
	"log"           // New import
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message, use
	// the http.Error() function to send an Internal Server Error response to the
	// user, and then return from the handler so no subsequent code is executed.
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Then we use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func listSnippets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly list snippets"))
}

func newSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly form for create snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create snippet by ID"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display snippet by ID"))
}

func editSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for edit snippet"))
}

func updateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update snippet by ID"))
}

func deleteSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete snippet by ID"))
}
