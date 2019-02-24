package main

import (
	"net/http"
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

func postHandler(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
}
