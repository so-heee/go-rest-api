package service

import (
	"github.com/so-heee/go-rest-api/api/domain/model"
	"github.com/so-heee/go-rest-api/api/usecase/repository"
)

// UserService.
type UserService struct {
	UserRepository repository.UserRepository
}

func (s *UserService) UserById(id int) (user model.User, err error) {
	user, err = s.UserRepository.FindById(id)
	return
}
