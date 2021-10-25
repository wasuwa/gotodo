package main

import (
	"database/sql"
	"gotodo/logging"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Task struct {
	Id       int
	Title    string
	Describe string
}

func main() {
	logging.LoggingSettings("log/development.log")
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func dbConn() (db *sql.DB) {
	connStr := "user=suwayouta dbname=todo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

var templates = template.Must(template.ParseGlob("view/*"))

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	sql := "SELECT * FROM tasks ORDER BY id DESC"
	rows, err := db.Query(sql)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalln(err)
	}
	var ts []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Title, &t.Describe)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln(err)
		}
		ts = append(ts, t)
	}
	templates.ExecuteTemplate(w, "index.html", ts)
}

func create(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
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
