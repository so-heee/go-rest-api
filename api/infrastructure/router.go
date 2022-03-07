package infrastructure

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/so-heee/go-rest-api/api/interfaces/controllers"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
)

type Handler struct {
	GetUserHandler
}

type GetUserHandler struct{}

func (h *GetUserHandler) GetUsersUserId(c echo.Context) error {
	return c.JSON(http.StatusOK, &oapi.GetUsersUserIdResponse{})
}

func Run() {
	// Echo instance
	e := echo.New()

	conf := NewSQLConfig("user", "password", "db", "sample", 3306)
	h, err := NewSQLHandler(&conf)
	if err != nil {
		log.Fatal(err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// swagger, err := oapi.GetSwagger()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// e.Use(oapimiddleware.OapiRequestValidator(swagger))

	userController := controllers.NewUserController(h)

	e.GET("/", health)
	// e.GET("/user/:id", userController.Show)

	oapi.RegisterHandlers(e, userController)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "health check")
}
