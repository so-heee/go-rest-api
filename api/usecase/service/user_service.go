package service

import (
	"github.com/so-heee/go-rest-api/api/domain/model"
	"github.com/so-heee/go-rest-api/api/usecase/repository"
)

// UserService.
type UserService struct {
	UserRepository repository.UserRepository
}

func (s *UserService) CreateUser(u *model.User) (user *model.User, err error) {
	user, err = s.UserRepository.CreateUser(u)
	return
}

func (s *UserService) UpdateUser(u *model.User) (user *model.User, err error) {
	user, err = s.UserRepository.UpdateUser(u)
	return
}

func (s *UserService) Users() (user []model.User, err error) {
	user, err = s.UserRepository.Users()
	return
}

func (s *UserService) UserById(id int) (user *model.User, err error) {
	user, err = s.UserRepository.FindById(id)
	return
}

func (s *UserService) UserByName(n string) (user *model.User, err error) {
	user, err = s.UserRepository.FindByName(n)
	return
}
