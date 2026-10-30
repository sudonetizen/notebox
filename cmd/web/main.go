package main

import (
	"log"
	"flag"
	"net/http"
) 

func main() {
	addr := flag.String("addr", ":4000", "http network address")
	flag.Parse()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /note/view/{id}", noteView)
	mux.HandleFunc("GET /note/create", noteCreate)
	mux.HandleFunc("POST /note/create", noteCreatePost)

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

