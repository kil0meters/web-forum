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
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type Comment struct {
	Date          time.Time
	Author        Account
	UpvoteCount   int32
	DownvoteCount int32
}

func main() {
	logPath := "server.log"
	httpPort := 8000

	log.Printf("logging to \"%s\"", logPath)

	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file %v", err)
	}

	logWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logWriter)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Printf("starting web server at http://localhost:%d", httpPort)

	r := mux.NewRouter()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", frontPageHandler)
	r.HandleFunc("/posts/{id}", postHandler)

	r.HandleFunc("/api/front/", frontPageApiHandler)

	err = http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r)
	if err != nil {
		log.Fatal(err)
	}
}
