package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/so-heee/go-rest-api/api/interfaces/database"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
	"github.com/so-heee/go-rest-api/api/usecase/service"
)

// UserController.
type UserController struct {
	UserService service.UserService
}

func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		UserService: service.UserService{
			UserRepository: &database.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) GetUsersUserId(c echo.Context, id int) (err error) {
	user, err := controller.UserService.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	dto := oapi.User{
		Id:   int64(user.Id),
		Name: &user.Name,
	}
	c.JSON(200, dto)
	return
}
