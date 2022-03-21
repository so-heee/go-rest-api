package repository

import "github.com/so-heee/go-rest-api/api/domain/model"

type UserRepository interface {
	CreateUser(u *model.User) (*model.User, error)
    UpdateUser(u *model.User) (*model.User, error)
	FindById(id int) (*model.User, error)
}
