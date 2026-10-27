package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"html/template"
) 


func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{"./ui/html/base.tmpl", "./ui/html/pages/home.tmpl"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return 
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func noteView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}	

	fmt.Fprintf(w, "display some kind of note with id %d \n", id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display form for creating note\n"))
}

func noteCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("save a note\n"))
}

