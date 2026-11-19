package main

import (
	"fmt"
	"errors"
	"strconv"
	"net/http"
	"notebox/internal/models"
) 


func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	notes, err := app.notes.Latest() 
	if err != nil {
		app.serverError(w, r, err)
		return 
	}

	app.render(w, r, http.StatusOK, "home.tmpl", templateData{Notes: notes})
}

func (app *application) noteView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}	

	note, err := app.notes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return 
	}

	app.render(w, r, http.StatusOK, "view.tmpl", templateData{Note: note})
}

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display form for creating note\n"))
}

func (app *application) noteCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "somekind of title"
	content := "somekind of content"
	expires := 4 

	id, err := app.notes.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return 
	}

	http.Redirect(w, r, fmt.Sprintf("/note/view/%d", id), http.StatusSeeOther)
}

