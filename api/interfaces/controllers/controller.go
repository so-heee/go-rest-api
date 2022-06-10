package controllers

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

	t, err := security.GenerateAccessToken(user.Id)
	if err != nil {
		return convertError(err)
	}
	r, err := security.GenerateRefreshToken(user.Id)
	if err != nil {
		return convertError(err)
	}

	dto := oapi.AuthenticationResponse{
		AccessToken:  t,
		RefreshToken: r,
		TokenType:    "bearer",
	}
	c.JSON(http.StatusOK, dto)
	return
}

func (controller *Controller) RefreshAccessToken(c echo.Context, params oapi.RefreshAccessTokenParams) (err error) {
	u, err := security.VerifyRefreshToken(params.RefreshToken)
	if err != nil {
		return convertError(err)
	}

	i, err := strconv.Atoi(u)
	if err != nil {
		return convertError(err)
	}
	t, err := security.GenerateAccessToken(i)
	if err != nil {
		return convertError(err)
	}

	dto := oapi.AuthenticationResponse{
		AccessToken:  t,
		RefreshToken: params.RefreshToken,
		TokenType:    "bearer",
	}
	c.JSON(http.StatusOK, dto)
	return
}

func (controller *Controller) PostUser(c echo.Context) (err error) {
	var p oapi.UserPostRequest
	if err := c.Bind(&p); err != nil {
		return convertError(err)
	}
	if err := c.Validate(p); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			c.Logger().Error(k + ": " + err.Error())
		}
		return err
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
		Mail: &user.Mail,
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
	log.Info(*p.Name)
	log.Info(p.Mail)
	u := model.User{Id: id, Name: *p.Name, Mail: *p.Mail}
	log.Info("test2")
	user, err := controller.UserService.UpdateUser(&u)
	log.Info("test3")
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
