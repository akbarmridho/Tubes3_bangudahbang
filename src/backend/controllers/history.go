package controllers

import (
	"backend/configs"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllHistoryHandler(c echo.Context) error {
	response := models.Response[[]models.History]{}

	db := configs.DB.GetConnection()
	history := []models.History{}
	if err := db.Table("histories").
		Joins("INNER JOIN (SELECT session_id, MAX(created_at) AS latest_created_at FROM histories GROUP BY session_id) sub ON histories.session_id = sub.session_id AND histories.created_at = sub.latest_created_at").
		Find(&history).
		Error; err != nil {
		response.Message = "ERROR: FAILED TO GET HISTORY"
		return c.JSON(http.StatusBadRequest, response)
	}

	// if err := db.Find(&history).Error; err != nil {
	// 	response.Message = "ERROR: FAILED TO GET HISTORY"
	// 	return c.JSON(http.StatusBadRequest, response)
	// }

	response.Message = "Success"
	response.Data = history

	return c.JSON(http.StatusOK, response)
}

func GetHistoryHandler(c echo.Context) error {
	response := models.Response[[]models.History]{}

	db := configs.DB.GetConnection()
	history := []models.History{}

	if err := db.Where("session_id = ?", c.Param("id")).Find(&history).Error; err != nil {
		response.Message = "ERROR: FAILED TO GET HISTORY"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Success"
	response.Data = history

	return c.JSON(http.StatusOK, response)
}
