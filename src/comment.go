package web_forum

import "time"

type Comment struct {
	Date          time.Time `json:"date"`
	Author        Account   `json:"author"`
	UpvoteCount   int32     `json:"upvote_count"`
	DownvoteCount int32     `json:"downvote_count"`
}
