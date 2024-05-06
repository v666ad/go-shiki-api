package types

type Me struct {
	ID         uint       `json:"id"`
	Nickname   string     `json:"nickname"`
	Avatar     string     `json:"avatar"`
	Image      UserAvatar `json:"image"`
	LastOnline string     `json:"last_online_at"`
	URL        string     `json:"url"`
	Name       string     `json:"name,omitempty"`
	Sex        string     `json:"sex,omitempty"`
	Website    string     `json:"website,omitempty"`
	BirthOn    string     `json:"birth_on,omitempty"`
	FullYears  int        `json:"full_years,omitempty"`
	Locale     string     `json:"locale"`
}
