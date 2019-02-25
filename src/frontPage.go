package web_forum

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type FrontPage struct {
	Posts      []Post `json:"posts"`
	PageNumber int32  `json:"page_number"`
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
			Author: Account{
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

	// for i := 0; i < 10; i++ {
	// 	posts[i] = Post{
	// 		Id:    randomId(),
	// 		Title: "Lorem Ipsum",
	// 		Date:  time.Now(),
	// 		Author: Account{
	// 			Username:     "kilometers",
	// 			CreatedAt:    time.Now(),
	// 			PasswordHash: "ajsdfj",
	// 		},
	// 		Body:          "{}",
	// 		UpvoteCount:   int32(i),
	// 		DownvoteCount: 0,
	// 		CommentCount:  0,
	// 	}
	// }
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
