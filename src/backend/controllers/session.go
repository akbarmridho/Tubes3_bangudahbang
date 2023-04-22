package controllers

import (
	"backend/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type SessionResponse struct {
	SessionId string `json:"session_id"`
}

func GetSessionHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Response[SessionResponse]{
		Message: "Ok",
		Data:    SessionResponse{SessionId: uuid.NewV4().String()},
	})
}
