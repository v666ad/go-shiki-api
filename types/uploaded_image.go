package types

type UploadedImage struct {
	ID      uint   `json:"id"`
	Preview string `json:"preview"`
	URL     string `json:"url"`
	BBCode  string `json:"bbcode"`
}
