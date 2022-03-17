package model

import "time"

// User.
type User struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Mail      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string, mail string, password string) *User {
	return &User{Name: name, Mail: mail, Password: password}
}
