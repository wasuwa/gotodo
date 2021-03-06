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

func (t *Task) Find(id string) error {
	db := database.DbConn()
	sql := "SELECT * FROM tasks WHERE id = $1"
	row := db.QueryRow(sql, id)
	err := row.Scan(&t.Id, &t.Title, &t.Describe)
	return err
}

func (t *Task) Update(id, ttl, d string) error {
	db := database.DbConn()
	sql := "UPDATE tasks SET title = $1,describe = $2 WHERE id = $3"
	_, err := db.Exec(sql, ttl, d, id)
	return err
}

func (t *Task) Destroy(id string) error {
	db := database.DbConn()
	sql := "DELETE FROM tasks WHERE id = $1"
	_, err := db.Exec(sql, id)
	return err
}
