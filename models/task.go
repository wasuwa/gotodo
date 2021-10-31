package models

import (
	"database/sql"
	"gotodo/database"

	_ "github.com/lib/pq"
)

type Task struct {
	Id       int
	Title    string
	Describe string
}

func (t *Task) All() ([]Task, error) {
	db := database.DbConn()
	sql := "SELECT * FROM tasks ORDER BY id DESC"
	rows, err := db.Query(sql)
	var ts []Task
	if err != nil {
		return ts, err
	}
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Describe)
		if err != nil {
			return ts, err
		}
		ts = append(ts, *t)
	}
	return ts, err
}
