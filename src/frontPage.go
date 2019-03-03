package web_forum

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type FrontPage struct {
	Posts      []Post `json:"posts"`
	PageNumber int32  `json:"page_number"`
}

func getFrontPageData() FrontPage {
	log.Print("Loading front page")

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil { // TODO: don't kill entire website upon single db error
		log.Fatal(err)
	}

	var (
		id            string
		title         string
		date          string
		author        string
		body          string
		upvoteCount   int32
		downvoteCount int32
		commentCount  int32
	)

	var posts []Post

	for rows.Next() {
		err := rows.Scan(&id, &title, &date, &author, &body,
			&upvoteCount, &downvoteCount, &commentCount)
		if err != nil { // TODO
			log.Fatal(err)
		}

		dateString, _ := time.Parse(time.RFC3339, date)

		posts = append(posts, Post{
			Id:    id,
			Title: title,
			Date:  dateString,
			Author: User{
				Username:     author,
				CreatedAt:    time.Now(),
				PasswordHash: "nil",
			},
			Body:          body,
			UpvoteCount:   upvoteCount,
			DownvoteCount: downvoteCount,
			CommentCount:  commentCount,
		})
	}

	rows.Close()

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
