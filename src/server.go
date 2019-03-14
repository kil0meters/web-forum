package web_forum

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartServer() {
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

	r.HandleFunc("/signup", signupPageHandler)
	r.HandleFunc("/signup/", signupPageHandler)

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
