package infrastructure

import (
	"log"
	"net/http"

	// oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/so-heee/go-rest-api/api/interfaces/controllers"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
)

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
	e.Use(middleware.CORS())

	g := e.Group("/v1")
	// swagger, err := oapi.GetSwagger()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// e.Use(oapimiddleware.OapiRequestValidator(swagger))

	controller := controllers.NewController(h)

	g.GET("/", health)
	oapi.RegisterHandlers(g, controller)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "health check")
}
