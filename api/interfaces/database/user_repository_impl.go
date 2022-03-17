package database

import "github.com/so-heee/go-rest-api/api/domain/model"

type UserRepository struct {
	SQLHandler
}

func (r *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	if err := r.Create(&u).Error(); err != nil {
		return &model.User{}, err
	}
	return u, nil
}

func (r *UserRepository) FindById(id int) (user *model.User, err error) {
	if err = r.First(&user, id).Error(); err != nil {
		return
	}
	return
}
