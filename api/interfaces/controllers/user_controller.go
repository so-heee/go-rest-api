package controllers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/so-heee/go-rest-api/api/interfaces/database"
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

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.UserService.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
