package controller

import (
	"log/slog"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nghuyentrangvt98/banhmioi/model"
	"github.com/nghuyentrangvt98/banhmioi/repo"
)

func GetProduct(c echo.Context) error {
	productID := c.Param("id")
	pID, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(400, model.Exception{Message: "invalid product ID"})
		return nil
	}
	product := repo.GetProduct(c.Request().Context(), pID)
	if product == nil {
		c.JSON(400, model.Exception{Message: "user not found"})
		return nil
	}
	c.JSON(200, product)
	return nil
}

func GetMultiProduct(c echo.Context) error {
	var productQuery model.ProductQuery
	err := c.Bind(&productQuery)
	if err != nil {
		slog.Info(err.Error())
	}
	products := repo.GetMultiProduct(c.Request().Context(), productQuery)
	c.JSON(200, products)
	return nil
}
