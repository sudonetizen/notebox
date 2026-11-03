package main

import "net/http"


func (app *application) routes() *http.ServeMux {
	// server mux and file server 
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	
	// handle and handle functions 
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /note/view/{id}", app.noteView)
	mux.HandleFunc("GET /note/create", app.noteCreate)
	mux.HandleFunc("POST /note/create", app.noteCreatePost)

	return mux 
}
