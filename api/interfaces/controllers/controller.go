package controllers

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
	req := &oapi.AuthenticationRequest{}
	err = c.Bind(req)
	if err != nil {
		return
	}

	claims := &security.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	dto := oapi.AuthenticationResponse{
		Token: t,
	}
	c.JSON(200, dto)
	return
}

func (controller *Controller) GetUserByID(c echo.Context, id int) (err error) {
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

func (controller *Controller) GetTweetByID(c echo.Context, id int) (err error) {
	tweet, err := controller.TweetService.TweetById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	dto := oapi.Tweet{
		Id:     int64(tweet.Id),
		Text:   &tweet.Text,
		UserId: &tweet.UserId,
	}
	c.JSON(200, dto)
	return
}
