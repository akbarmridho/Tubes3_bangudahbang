package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/services"
	"net/http"
	"regexp"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type QueryRequest struct {
	SessionId uuid.UUID `json:"session_id" validate:"required,uuid4"`
	Input     string    `json:"input" validate:"required"`
	IsKMP     bool      `json:"is_kmp" validate:"required"`
}

func GetQueryHandler(c echo.Context) error {
	queryRequest := new(QueryRequest)
	response := models.Response[models.History]{}

	if err := c.Bind(queryRequest); err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(queryRequest); err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	// calculator regex
	var onlyMathRegex = regexp.MustCompile(`^[\s\d()+\-*/^]+$`)
	// referensi https://www.regular-expressions.info/dates.html
	// todo ubah agar bisa mencakup lebih banyak format
	var dateRegex = regexp.MustCompile(`^(19|20)\d\d[- -.](0[1-9]|1[012])[- -.](0[1-9]|[12][0-9]|3[01])$`)
	var addQueryRegex = regexp.MustCompile(`^[Tt]ambahkan pertanyaan (.*) dengan jawaban (.*)$`)
	var deleteQueryRegex = regexp.MustCompile(`^[Hh]apus pertanyaan (.*)$`)

	var message string
	var err error

	if dateRegex.MatchString(queryRequest.Input) {
		message, err = services.GetDay(queryRequest.Input)
	} else if onlyMathRegex.MatchString(queryRequest.Input) {
		message, err = services.Calculate(queryRequest.Input)
	} else if addQueryRegex.MatchString(queryRequest.Input) {
		matches := addQueryRegex.FindStringSubmatch(queryRequest.Input)
		message, err = services.AddQuery(matches[1], matches[2])
	} else if deleteQueryRegex.MatchString(queryRequest.Input) {
		matches := deleteQueryRegex.FindStringSubmatch(queryRequest.Input)
		message, err = services.DeleteQuery(matches[1])
	} else {
		message, err = services.MatchQuery(queryRequest.Input, queryRequest.IsKMP)
	}

	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	history := models.History{
		SessionId: queryRequest.SessionId,
		UserQuery: queryRequest.Input,
		Response:  message,
	}

	db := configs.DB.GetConnection()

	if err := db.Create(&history).Error; err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "Ok"
	response.Data = history
	return c.JSON(http.StatusOK, response)
}
