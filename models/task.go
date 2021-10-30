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
	rows, _ := db.Query(sql)
	ts, err := t.rowsToConvertSlice(rows)
	if err != nil {
		return nil, err
	}
	return ts, err
}

func (t *Task) rowsToConvertSlice(rows *sql.Rows) ([]Task, error) {
	var ts []Task
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Describe)
		if err != nil {
			return nil, err
		}
		ts = append(ts, *t)
	}
	return ts, nil
}
