package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"github.com/eastwood/got-it-done/handlers"
)

var PORT = ":8080";

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "I am healthy!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("./public/index.html");
  if err != nil {
    panic("Could not load index.html page")
  }
  t.Execute(w, "base");
}

func main() {
  http.HandleFunc("/api/health-check", healthCheckHandler)
  http.HandleFunc("/", indexHandler)

  http.HandleFunc("/api/todo", handlers.TodoHandler);
  http.HandleFunc("/api/todos", handlers.ListHandler);

  fmt.Printf("Listening on port: %s", PORT)
  log.Fatal(http.ListenAndServe(PORT, nil))
}
