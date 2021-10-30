package main

import (
	"gotodo/controllers"
	"gotodo/logging"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	logging.LoggingSettings("log/development.log")
	http.HandleFunc("/", controllers.TaskIndex)
	http.HandleFunc("/create", controllers.TaskCreate)
	http.HandleFunc("/edit", controllers.TaskEdit)
	http.HandleFunc("/update", controllers.TaskUpdate)
	http.HandleFunc("/destroy", controllers.TaskDestroy)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
