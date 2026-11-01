package main

import (
	"os"
	"flag"
	"log/slog"
	"net/http"
) 


// application struct 
type application struct {
	logger *slog.Logger 
}


func main() {
	// flags
	addr := flag.String("addr", ":4000", "http network address")
	flag.Parse()

	
	// logger 
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))


	// int application struct 
	app := &application{
		logger: logger,
	}


	// server mux and file server 
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	
	// handle and handle functions 
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /note/view/{id}", app.noteView)
	mux.HandleFunc("GET /note/create", app.noteCreate)
	mux.HandleFunc("POST /note/create", app.noteCreatePost)


	logger.Info("starting server", "addr", *addr)


	// listen and serve 
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

