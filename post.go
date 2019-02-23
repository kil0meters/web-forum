package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Post struct {
	Id            string
	Title         string
	Date          time.Time
	Author        Account
	Body          string `json:"body"`
	UpvoteCount   int32
	DownvoteCount int32
	CommentCount  int32
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
}
