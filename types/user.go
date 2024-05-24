package types

import "time"

type User struct {
	ID           uint       `json:"id"`
	Nickname     string     `json:"nickname"`
	Avatar       string     `json:"avatar"`
	Image        UserAvatar `json:"image"`
	LastOnlineAt time.Time  `json:"last_online_at"`
	URL          string     `json:"url"`
}
