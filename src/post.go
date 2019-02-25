package web_forum

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
	// "github.com/gorilla/mux"
)

type Post struct {
	Id            string    `json:"id"`
	Title         string    `json:"title"`
	Date          time.Time `json:"date"`
	Author        Account   `json:"author"`
	Body          string    `json:"body"`
	UpvoteCount   int32     `json:"upvote_count"`
	DownvoteCount int32     `json:"downvote_count"`
	CommentCount  int32     `json:"comment_count"`
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/postEditor.html")
	if err != nil {
		fmt.Fprintf(w, "503: ", err)
	} else {
		tmpl.Execute(w, nil)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
}
