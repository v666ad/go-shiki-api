package types

type Character struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name"`
	Russian           *string `json:"russian"`
	Image             Image   `json:"image"`
	URL               string  `json:"url"`
	Altname           *string `json:"altname"`
	Japanese          string  `json:"japanese"`
	Description       *string `json:"description"`
	DescriptionSource *string `json:"description_source"`
	ThreadID          uint    `json:"thread_id"`
	TopicID           uint    `json:"topic_id"`
}
