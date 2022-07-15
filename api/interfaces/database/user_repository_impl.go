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

func (r *UserRepository) UpdateUser(u *model.User) (*model.User, error) {
	if err := r.Updates(&u).Error(); err != nil {
		return &model.User{}, err
	}
	return u, nil
}

func (r *UserRepository) Users() (users []model.User, err error) {
	if err = r.Find(&users).Error(); err != nil {
		return
	}
	return
}

func (r *UserRepository) FindById(id int) (user *model.User, err error) {
	if err = r.First(&user, id).Error(); err != nil {
		return
	}
	return
}

func (r *UserRepository) FindByName(n string) (user *model.User, err error) {
	if err = r.Where("name = ?", n).First(&user).Error(); err != nil {
		return
	}
	return
}
