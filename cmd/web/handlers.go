package main

import (
	"fmt"
	"errors"
	"strconv"
	"net/http"
	"html/template"
	"notebox/internal/models"
) 


func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{"./ui/html/base.tmpl", "./ui/html/partials/nav.tmpl", "./ui/html/pages/home.tmpl"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return 
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}

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

	files := []string{"./ui/html/base.tmpl", "./ui/html/partials/nav.tmpl", "./ui/html/pages/view.tmpl"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return 
	}

	err = ts.ExecuteTemplate(w, "base", note)
	if err != nil {
		app.serverError(w, r, err) 
	}
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

