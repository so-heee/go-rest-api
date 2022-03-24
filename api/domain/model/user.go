package model

import (
	"time"

	"github.com/so-heee/go-rest-api/api/domain/value"
)

// User.
type User struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Mail      string
	Password  value.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string, mail string, password string) *User {
	return &User{Name: name, Mail: mail, Password: value.Password(password)}
}
