package infrastructure

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/so-heee/go-rest-api/api/interfaces/controllers"
)

func Run() {
	// Echo instance
	e := echo.New()

	conf := NewSQLConfig("user", "password", "db", "sample", 3306)
	h, err := NewSQLHandler(&conf)
	if err != nil {
		log.Fatal(err)
	}

	userController := controllers.NewUserController(h)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する
	e.GET("/user/:id", func(c echo.Context) error { return userController.Show(c) })

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World2!!")
}
