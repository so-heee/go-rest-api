package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/so-heee/go-rest-api/api/domain/model"
	"github.com/so-heee/go-rest-api/api/domain/value"
	"github.com/so-heee/go-rest-api/api/infrastructure/security"
	"github.com/so-heee/go-rest-api/api/interfaces/database"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
	"github.com/so-heee/go-rest-api/api/usecase/service"
)

// Controller.
type Controller struct {
	UserService  service.UserService
	TweetService service.TweetService
}

func NewController(sqlHandler database.SQLHandler) *Controller {
	return &Controller{
		UserService: service.UserService{
			UserRepository: &database.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
		TweetService: service.TweetService{
			TweetRepository: &database.TweetRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *Controller) Authenticate(c echo.Context) (err error) {
	name := c.FormValue("name")
	password := c.FormValue("password")
	req := oapi.AuthenticationRequest{
		Name:     name,
		Password: password,
	}
	user, err := controller.UserService.UserByName(req.Name)
	if err != nil {
		return convertError(err)
	}
	err = user.Password.Verify(req.Password)
	if err != nil {
		return convertError(err)
	}

	t, err := security.GenerateToken()
	if err != nil {
		return err
	}

	dto := oapi.AuthenticationResponse{
		AccessToken: t,
	}
	c.JSON(http.StatusOK, dto)
	return
}

func (controller *Controller) PostUser(c echo.Context) (err error) {
	var p oapi.UserPostRequest
	if err := c.Bind(&p); err != nil {
		return convertError(err)
	}
	u := model.NewUser(p.Name, p.Mail, p.Password)
	x, err := u.Password.ConvertHash()
	if err != nil {
		return convertError(err)
	}
	u.Password = value.Password(x)

	user, err := controller.UserService.CreateUser(u)
	if err != nil {
		return convertError(err)
	}

	dto := oapi.User{
		Id:   int64(user.Id),
		Name: &user.Name,
		Mail: &user.Mail,
	}
	c.JSON(http.StatusCreated, dto)
	return
}

func (controller *Controller) GetUserByID(c echo.Context, id int) (err error) {
	user, err := controller.UserService.UserById(id)
	if err != nil {
		return convertError(err)
	}
	dto := oapi.User{
		Id:   int64(user.Id),
		Name: &user.Name,
	}
	c.JSON(http.StatusOK, dto)
	return
}

func (controller *Controller) PatchUser(c echo.Context, id int) (err error) {
	var p oapi.UserPatchRequest
	if err := c.Bind(&p); err != nil {
		return convertError(err)
	}
	u := model.User{Id: id, Name: p.Name}
	user, err := controller.UserService.UpdateUser(&u)
	if err != nil {
		return convertError(err)
	}

	dto := oapi.User{
		Id:   int64(user.Id),
		Name: &user.Name,
		Mail: &user.Mail,
	}
	c.JSON(http.StatusOK, dto)
	return
}

func (controller *Controller) GetTweetByID(c echo.Context, id int) (err error) {
	tweet, err := controller.TweetService.TweetById(id)
	if err != nil {
		return convertError(err)
	}
	dto := oapi.Tweet{
		Id:     int64(tweet.Id),
		Text:   &tweet.Text,
		UserId: &tweet.UserId,
	}
	c.JSON(http.StatusOK, dto)
	return
}
