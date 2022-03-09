package infrastructure

import (
	"log"
	"net/http"
	"time"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/so-heee/go-rest-api/api/interfaces/controllers"
	oapi "github.com/so-heee/go-rest-api/api/openapi"
)

func Run() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	g := e.Group("/v1")

	swagger, err := oapi.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	// jwtConfig := middleware.JWTConfig{
	// 	Skipper: func(e echo.Context) bool {
	// 		jwt, ok := e.Get("jwt").(bool)
	// 		if ok && jwt {
	// 			return false
	// 		}
	// 		return true
	// 	},
	// 	Claims:     &JWTCustomClaims{},
	// 	SigningKey: []byte("SECRET_KEY"), // SDK_AUTH_TOKEN change on next deploy
	// }
	// g.Use(middleware.JWTWithConfig(jwtConfig))
	g.Use(middleware.JWT([]byte("secret")))

	g.Use(oapimiddleware.OapiRequestValidator(swagger))

	conf := NewSQLConfig("user", "password", "db", "sample", 3306)
	h, err := NewSQLHandler(&conf)
	if err != nil {
		log.Fatal(err)
	}

	e.GET("/", healthCheck)
	e.POST("/login", login)
	oapi.RegisterHandlers(g, controllers.NewController(h))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// とりあえずのパスワード認証
	if username != "taro" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// トークン作成
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Taro"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
