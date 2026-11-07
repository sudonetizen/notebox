package models 

import (
	"database/sql"
	"time"
)


type Note struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time 
}


type NoteModel struct {
	DB *sql.DB
}


func (m *NoteModel) Insert(title string, content string, expires int) (int, error) {
	stmt := "insert into notes (title, content, created, expires) values (?, ?, utc_timestamp(), date_add(utc_timestamp(), interval ? day))"

	rst, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := rst.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil 
}


func (m *NoteModel) Get(id int) (Note, error) {
	return Note{}, nil 
}


func (m *NoteModel) Latest() ([]Note, error) {
	return nil, nil 
}
