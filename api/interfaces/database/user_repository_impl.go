package database

import "github.com/so-heee/go-rest-api/api/domain/model"

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) FindById(id int) (user model.User, err error) {
	if err = repo.Find(&user, id).Error(); err != nil {
		return
	}
	return
}
