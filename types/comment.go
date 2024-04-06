package types

type Comment struct {
	ID              uint   `json:"id"`
	UserID          uint   `json:"user_id"`
	CommentableID   uint   `json:"commentable_id"`
	CommentableType string `json:"commentable_type"`
	Body            string `json:"body"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	IsOfftopic      bool   `json:"is_offtopic"`
	User            User   `json:"user"`
}
