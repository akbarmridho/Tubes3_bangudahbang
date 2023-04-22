package api

import (
	"backend/configs"
	"backend/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Validator = &configs.RequestValidator{Validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.HistoryRoute(e)
	routes.SessionRoute(e)
	routes.QueryRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
