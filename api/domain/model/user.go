package model

import (
	"time"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/so-heee/go-rest-api/api/domain/value"
)

// User.
type User struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Mail      types.Email
	Password  value.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string, mail types.Email, password string) *User {
	return &User{Name: name, Mail: mail, Password: value.Password(password)}
}
