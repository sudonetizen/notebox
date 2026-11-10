package models 

import (
	"time"
	"errors"
	"database/sql"
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
	stmt := "select id, title, content, created, expires from notes where expires > utc_timestamp() and id = ?"
	row := m.DB.QueryRow(stmt, id)
	n := Note{} 
	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Note{}, ErrNoRecord
		} else {
			return Note{}, err 
		}
	}

	return n, nil 
}


func (m *NoteModel) Latest() ([]Note, error) {
	stmt := "select id, title, content, created, expires from notes where expires > utc_timestamp() order by id desc limit 10"
	rows, err := m.DB.Query(stmt)	
	if err != nil {
		return nil, err 
	}
	defer rows.Close() 
	
	notes := []Note{} 
	for rows.Next() {
		n := Note{} 
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Expires)
		if err != nil {
			return nil, err 
		}
		notes = append(notes, n)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil 
}
