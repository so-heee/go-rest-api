package repository

import "github.com/so-heee/go-rest-api/api/domain/model"

type UserRepository interface {
	FindById(id int) (model.User, error)
}
