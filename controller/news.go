package controller

import (
	"log/slog"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nghuyentrangvt98/banhmioi/model"
	"github.com/nghuyentrangvt98/banhmioi/repo"
)

func GetNews(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, model.Exception{Message: "invalid news ID"})
		return nil
	}
	new := repo.GetNews(c.Request().Context(), ID)
	if new == nil {
		c.JSON(400, model.Exception{Message: "news not found"})
		return nil
	}
	c.JSON(200, new)
	return nil
}

func GetMultiNews(c echo.Context) error {
	var newsQuery model.NewsQuery
	err := c.Bind(&newsQuery)
	if err != nil {
		slog.Info(err.Error())
	}
	news := repo.GetMultiNews(c.Request().Context(), newsQuery)
	c.JSON(200, news)
	return nil
}
