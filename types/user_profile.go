package types

type UserProfile struct {
	ID           int64    `json:"id"`
	Nickname     string   `json:"nickname"`
	Avatar       string   `json:"avatar"`
	Image        Image    `json:"image"`
	LastOnlineAt string   `json:"last_online_at"`
	URL          string   `json:"url"`
	Name         *string  `json:"name,omitempty"`
	Sex          *string  `json:"sex,omitempty"`
	FullYears    *int     `json:"full_years"`
	LastOnline   string   `json:"last_online"`
	Website      string   `json:"website"`
	Location     *string  `json:"location,omitempty"`
	Banned       bool     `json:"banned"`
	About        string   `json:"about"`
	AboutHTML    string   `json:"about_html"`
	CommonInfo   []string `json:"common_info"`
	ShowComments bool     `json:"show_comments"`
	InFriends    *bool    `json:"in_friends"`
	IsIgnored    bool     `json:"is_ignored"`
	StyleID      *int     `json:"style_id"`
}
