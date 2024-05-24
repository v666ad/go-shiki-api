package types

import "time"

type TopicForum struct {
	ID        uint   `json:"id"`
	Position  uint   `json:"position"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	URL       string `json:"url"`
}

type Topic struct {
	ID            uint       `json:"id"`
	TopicTitle    string     `json:"topic_title"`
	Body          string     `json:"body"`
	HtmlBody      string     `json:"html_body"`
	HtmlFooter    string     `json:"html_footer"`
	CreatedAt     time.Time  `json:"created_at"`
	CommentsCount uint       `json:"comments_count"`
	Forum         TopicForum `json:"forum"`
	User          User       `json:"user"`
}
