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
	var t models.Task
	ts, err := t.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	templates.ExecuteTemplate(w, "index.html", ts)
}

func TaskCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var t models.Task
		ttl := r.FormValue("title")
		d := r.FormValue("describe")
		err := t.Create(ttl, d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func TaskEdit(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	id := r.URL.Query().Get("id")
	err := t.Find(id)
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
