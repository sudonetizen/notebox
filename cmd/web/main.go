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


	logger.Info("starting server", "addr", *addr)


	// listen and serve 
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

