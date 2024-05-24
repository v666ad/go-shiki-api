package types

import "time"

type Message struct {
	ID        uint64    `json:"id"`
	Kind      string    `json:"kind"`
	Read      bool      `json:"read"`
	Body      string    `json:"body"`
	HtmlBody  string    `json:"html_body"`
	CreatedAt time.Time `json:"created_at"`
	From      User      `json:"from"`
	To        User      `json:"to"`
}
