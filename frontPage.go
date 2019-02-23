package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type FrontPage struct {
	Posts      []Post
	PageNumber int32
}

// I should probably make sure two posts don't share the same id as that could
// break things
var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV1234567890@$^-_+"
var idLength = 6

func randomId() string {
	id := make([]byte, idLength)

	for i := 0; i < idLength; i++ {
		id[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(id)
}

func getFrontPageData() FrontPage {
	posts := make([]Post, 10)

	for i := 0; i < 10; i++ {
		posts[i] = Post{
			Id:    randomId(),
			Title: "Lorem Ipsum",
			Date:  time.Now(),
			Author: Account{
				Username:     "kilometers",
				CreatedAt:    time.Now(),
				PasswordHash: "ajsdfj",
			},
			Body:          "{}",
			UpvoteCount:   int32(i),
			DownvoteCount: 0,
			CommentCount:  0,
		}
	}
	return FrontPage{
		Posts:      posts,
		PageNumber: 1,
	}
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	frontPage := getFrontPageData()

	tmpl, err := template.ParseFiles("template/frontPage.html")

	if err != nil {
		fmt.Fprint(w, "503: ", err)
	} else {
		tmpl.Execute(w, frontPage)
	}
}

func frontPageApiHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(getFrontPageData())
	if err != nil {
		fmt.Fprint(w, "503: ", err)
	} else {
		fmt.Fprintf(w, "%s", string(jsonBytes))
	}
}
