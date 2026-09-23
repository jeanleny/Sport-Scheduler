package main

import (
	sql "database/sql"
	//"io"
	_ "github.com/jackc/pgx/v5/stdlib"
	"fmt"
	"log"
	"net/http"
	frontend "sportS/frontend"
)

type s_db struct {
	db *sql.DB
}

func connectDB () (*s_db, error) {
	db, err := sql.Open("pgx", "postgres://user:pass@db:5432/scheduler?sslmode=disable")
	database := &s_db{db : db}
	return database, err
}

func main() {
	db, err := connectDB()
	front := &frontend.S_frontEnd{}
	if err != nil {
		fmt.Println("failed to create database")
		fmt.Println(err)
	}
	router := http.NewServeMux()
	router.HandleFunc("/", front.Index)
	router.HandleFunc("/clicked", front.Clicked)
	fmt.Println("Server CREATION")
	log.Fatal(http.ListenAndServe(":8080", router))
	db.db.Close()
}
