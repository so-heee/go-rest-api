package model

import "time"

// Tweet.
type Tweet struct {
	Id        int
	Text      string
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
