package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Account struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

type Comment struct {
	Date          time.Time `json:"date"`
	Author        Account   `json:"author"`
	UpvoteCount   int32     `json:"upvote_count"`
	DownvoteCount int32     `json:"downvote_count"`
}

func main() {
	logPath := "server.log"
	httpPort := 8000

	log.Printf("logging to \"%s\"", logPath)

	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file %v", err)
	}

	initializeDB()

	logWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logWriter)
	log.SetFlags(log.Ldate | log.Ltime)

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", frontPageHandler)

	r.HandleFunc("/submit", newPostHandler)
	r.HandleFunc("/submit/", newPostHandler)

	r.HandleFunc("/posts/{id}", postHandler)

	r.HandleFunc("/api/front", frontPageApiHandler)
	r.HandleFunc("/api/front/", frontPageApiHandler)
	// r.HandleFunc("/api/submit")

	log.Printf("starting web server at http://localhost:%d", httpPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r)
	if err != nil {
		log.Fatal(err)
	}
}
