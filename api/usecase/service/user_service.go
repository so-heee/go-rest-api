package service

import (
	"github.com/so-heee/go-rest-api/api/domain/model"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
	"github.com/so-heee/go-rest-api/api/usecase/repository"
)

// UserService.
type UserService struct {
	UserRepository repository.UserRepository
}

func (s *UserService) CreateUser(u *model.User) (*model.User, error) {
	return s.UserRepository.CreateUser(u)
}

func (s *UserService) UpdateUser(u *model.User) (*model.User, error) {
	return s.UserRepository.UpdateUser(u)
}

func (s *UserService) Users() ([]oapi.User, error) {
	users, err := s.UserRepository.Users()
	if err != nil {
		return nil, err
	}
	dtos := []oapi.User{}
	for _, u := range users {
		mail := u.Mail
		name := u.Name
		dto := oapi.User{
			Id:   int64(u.Id),
			Mail: &mail,
			Name: &name,
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}

func (s *UserService) UserById(id int) (oapi.User, error) {
	u, err := s.UserRepository.FindById(id)
	if err != nil {
		return oapi.User{}, err
	}
	dto := oapi.User{
		Id:   int64(u.Id),
		Mail: &u.Mail,
		Name: &u.Name,
	}
	return dto, nil
}

func (s *UserService) UserByName(n string) (*model.User, error) {
	return s.UserRepository.FindByName(n)
}
