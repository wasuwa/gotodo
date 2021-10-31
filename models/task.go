package models

import (
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

func (t *Task) Create(ttl string, d string) error {
	db := database.DbConn()
	sql := "INSERT INTO tasks (title, describe) VALUES ($1, $2)"
	if ttl == "" {
		ttl = "タイトルなし"
	}
	if d == "" {
		d = "説明なし"
	}
	_, err := db.Exec(sql, ttl, d)
	return err
}
