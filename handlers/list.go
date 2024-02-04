package handlers

import (
	"fmt"
	"net/http"
  "database/sql"
  _ "modernc.org/sqlite"
)

var db, err = sql.Open("sqlite", "./db/todos.db")

func ListHandler(w http.ResponseWriter, r *http.Request) {
  if (err != nil) {
    panic("Could not open connection to database")
  }
  defer db.Close()

  sql := 
  `CREATE TABLE events (
    id INTEGER PRIMARY KEY,  
    title TEXT NOT NULL,
    description TEXT NOT NULL
  );`

  _, err = db.Exec(sql);
  if (err != nil) {
    panic("Could not execute SQL query")
  }
  fmt.Fprintf(w, "List Handler")
}
