package types

import "time"

type DialogMessage struct {
	ID        uint64    `json:"id"` // ~ 3 млрд
	Kind      string    `json:"kind"`
	Read      bool      `json:"read"`
	Body      string    `json:"body"`
	HtmlBody  string    `json:"html_body"`
	CreatedAt time.Time `json:"created_at"`
}

type Dialog struct {
	TargetUser User          `json:"target_user"`
	Message    DialogMessage `json:"message"`
}
