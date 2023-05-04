package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/services"
	"backend/utils"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"regexp"
)

type QueryRequest struct {
	SessionId string `json:"session_id" form:"session_id" query:"input" validate:"required,uuid4"`
	Input     string `json:"input" form:"input" query:"input" validate:"required"`
	IsKMP     bool   `json:"is_kmp" form:"is_kmp" query:"is_kmp" validate:"boolean"`
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

	queryRequest.Input = utils.CleanString(queryRequest.Input)

	// calculator regex
	var onlyMathRegex = regexp.MustCompile(`[\s\d()+\-*/^]+`)
	var dateRegex = regexp.MustCompile(`(\d{4})[- -.](0*[1-9]|1[012])[- -.]([12][0-9]|3[01]|0*[1-9])|([12][0-9]|3[01]|0*[1-9])[- -.](0*[1-9]|1[012])[- -.](\d{4})`)
	var addQueryRegex = regexp.MustCompile(`^[Tt]ambahkan pertanyaan (.*) dengan jawaban (.*)$`)
	var deleteQueryRegex = regexp.MustCompile(`^[Hh]apus pertanyaan (.*)$`)

	var message string
	var err error

	if dateRegex.MatchString(queryRequest.Input) {
		match := dateRegex.FindAllString(queryRequest.Input, 1)
		message, err = services.GetDay(utils.CleanString(match[0]))
	} else if onlyMathRegex.MatchString(queryRequest.Input) {
		match := onlyMathRegex.FindAllString(queryRequest.Input, 1)
		message, err = services.Calculate(utils.CleanString(match[0]))
	} else if addQueryRegex.MatchString(queryRequest.Input) {
		matches := addQueryRegex.FindStringSubmatch(queryRequest.Input)
		message, err = services.AddQuery(utils.CleanString(matches[1]), matches[2])
	} else if deleteQueryRegex.MatchString(queryRequest.Input) {
		matches := deleteQueryRegex.FindStringSubmatch(queryRequest.Input)
		message, err = services.DeleteQuery(utils.CleanString(matches[1]))
	} else {
		message, err = services.MatchQuery(queryRequest.Input, queryRequest.IsKMP)
	}

	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	sessionId, err := uuid.FromString(queryRequest.SessionId)

	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	history := models.History{
		SessionId: sessionId,
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
