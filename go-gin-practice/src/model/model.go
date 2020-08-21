package model

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Task struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json]"updated_at"`
	Title     string    `json:"title"`
}

func DBConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "todo"
	dbPass := os.Getenv("MYSQL_TODO_PASSWORD")
	dbName := "tododb"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
