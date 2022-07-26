package repository

import "github.com/so-heee/go-rest-api/api/domain/model"

type UserRepository interface {
	CreateUser(u *model.User) (*model.User, error)
	UpdateUser(u *model.User) (*model.User, error)
	Users() ([]model.User, error)
	FindById(id int) (model.User, error)
	FindByName(n string) (*model.User, error)
}
