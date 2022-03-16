package model

import "time"

// User.
type User struct {
	Id        int
	Name      string
	Mail      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
