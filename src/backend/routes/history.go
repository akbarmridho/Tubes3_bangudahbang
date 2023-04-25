package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func HistoryRoute(e *echo.Echo) {
	e.GET("/history", controllers.GetAllHistoryHandler)
	e.GET("/history/:id", controllers.GetHistoryHandler)
}
