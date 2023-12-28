package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tesk-task-betera/models"
	"tesk-task-betera/service"
)

type Apod struct {
	serv *service.Apod
}

func NewApodHandler(serv *service.Apod) *Apod {
	return &Apod{serv: serv}
}

func (handler *Apod) getApods(ctx echo.Context) error {
	var getRecordsRequest models.GetApodsRequest
	if err := ctx.Bind(&getRecordsRequest); err != nil {
		ctx.Logger().Error(err)
		return echo.ErrBadRequest
	}

	if getRecordsRequest.Limit == nil {
		var limit int32 = 1
		getRecordsRequest.Limit = &limit
	}

	if getRecordsRequest.Offset == nil {
		var offset int32 = 0
		getRecordsRequest.Offset = &offset
	}

	getApodsResponse, err := handler.serv.GetApods(ctx.Request().Context(), &getRecordsRequest)

	if err != nil {
		ctx.Logger().Error(err)
		return echo.ErrBadRequest
	}

	if getRecordsRequest.Limit != nil {
		getApodsResponse.Limit = *getRecordsRequest.Limit
	}

	if getRecordsRequest.Offset != nil {
		getApodsResponse.Offset = *getRecordsRequest.Offset
	}

	if getApodsResponse.Apods == nil {
		getApodsResponse.Apods = []models.ApodDto{}
	}

	return ctx.JSON(http.StatusOK, getApodsResponse)
}

func (handler *Apod) getApodsByDate(ctx echo.Context) error {
	var getApodByDateRequest models.GetApodByDateRequest
	dateString := ctx.Param("date")

	err := getApodByDateRequest.Date.UnmarshalJSON([]byte(dateString))
	if err != nil {
		ctx.Logger().Error(err)
	}

	getApodByDateResponse, err := handler.serv.GetApodsByDate(ctx.Request().Context(), getApodByDateRequest)
	if err != nil {
		ctx.Logger().Error(err)
		return echo.ErrBadRequest
	}

	return ctx.JSON(http.StatusOK, getApodByDateResponse)
}
