package main

import (
	"fmt"
	"net/http"
	"strconv"
) 


func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("1234\n"))
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

