package shikimori

import (
	"time"
)

var (
	ShikiTimeFormat = "2006-01-02T15:04:05.000Z07:00"
)

func GetTime(shikiTime string) (*time.Time, error) {
	t, err := time.Parse(shikiTime, ShikiTimeFormat)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
