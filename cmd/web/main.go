package main

import (
	"os"
	"flag"
	"log/slog"
	"net/http"
	"database/sql"
	"notebox/internal/models"

	"github.com/joho/godotenv" 
	_"github.com/go-sql-driver/mysql"
) 


// application struct 
type application struct {
	logger *slog.Logger 
	notes *models.NoteModel 
}


func main() {
	// vars 
	godotenv.Load()
	usr := os.Getenv("USR")
	pss := os.Getenv("PSS")


	// flags
	addr := flag.String("addr", ":4000", "http network address")
	dsn := flag.String("dsn", usr+":"+pss+"@/notebox?parseTime=true", "mysql data source name")
	flag.Parse()


	// logger 
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))


	// db
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()


	// int application struct 
	app := &application{
		logger: logger,
		notes: &models.NoteModel{DB: db}, 
	}


	logger.Info("starting server", "addr", *addr)


	// listen and serve 
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}


// openDB func
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err 
	}
	
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err 
	}
	
	return db, nil 
}

