package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//_, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
