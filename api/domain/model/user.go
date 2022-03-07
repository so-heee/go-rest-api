package model

import "time"

// User.
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
