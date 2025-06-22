package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// 1. Connect to MySQL
	dsn := "testDevelopment:fMpqR#253dG@tcp(127.0.0.1:3306)/fishingVideos?parseTime=true"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// 2. Set up routes
	http.HandleFunc("/health", healthHandler) // simple test endpoint

	// Add more handlers here:
	// http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/upload", uploadHandler)
	// http.HandleFunc("/videos", listVideosHandler)
	// etc...

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
