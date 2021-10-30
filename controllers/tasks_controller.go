package controllers

import (
	"gotodo/database"
	"gotodo/models"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("views/*"))

func TaskIndex(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	sql := "SELECT * FROM tasks ORDER BY id DESC"
	rows, err := db.Query(sql)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalln(err)
	}
	var ts []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.Id, &t.Title, &t.Describe)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln(err)
		}
		ts = append(ts, t)
	}
	templates.ExecuteTemplate(w, "index.html", ts)
}

func TaskCreate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	if r.Method == "POST" {
		sql := "INSERT INTO tasks (title, describe) VALUES ($1, $2)"
		t := r.FormValue("title")
		d := r.FormValue("describe")
		if t == "" {
			t = "タイトルなし"
		}
		if d == "" {
			d = "説明なし"
		}
		_, err := db.Exec(sql, t, d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func TaskEdit(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	sql := "SELECT * FROM tasks WHERE id = $1"
	id := r.URL.Query().Get("id")
	row := db.QueryRow(sql, id)
	var t models.Task
	err := row.Scan(&t.Id, &t.Title, &t.Describe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	templates.ExecuteTemplate(w, "edit.html", t)
}

func TaskUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	sql := "UPDATE tasks SET title = $1,describe = $2 WHERE id = $3"
	id := r.URL.Query().Get("id")
	t := r.FormValue("title")
	d := r.FormValue("describe")
	_, err := db.Exec(sql, t, d, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func TaskDestroy(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	sql := "DELETE FROM tasks WHERE id = $1"
	id := r.URL.Query().Get("id")
	_, err := db.Exec(sql, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalln(err)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
