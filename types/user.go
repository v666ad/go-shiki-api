package types

type User struct {
	ID           uint       `json:"id"`
	Nickname     string     `json:"nickname"`
	Avatar       string     `json:"avatar"`
	Image        UserAvatar `json:"image"`
	LastOnlineAt string     `json:"last_online_at"`
	URL          string     `json:"url"`
}
