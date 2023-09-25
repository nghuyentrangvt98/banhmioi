package controller

import (
	"log/slog"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nghuyentrangvt98/banhmioi/model"
	"github.com/nghuyentrangvt98/banhmioi/repo"
)

func CreateOrder(c echo.Context) error {
	var orderModel model.OrderCreate
	err := c.Bind(&orderModel)
	if err != nil {
		c.JSON(400, model.Exception{Message: err.Error()})
		return nil
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id, err := claims.GetSubject()
	if err != nil {
		slog.Error(err.Error())
	}
	idStr, _ := strconv.Atoi(id)
	cart := repo.CreateOrder(c.Request().Context(), orderModel, idStr)
	c.JSON(200, cart)
	return nil
}

func GetOrders(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id, err := claims.GetSubject()
	if err != nil {
		slog.Error(err.Error())
	}
	idStr, _ := strconv.Atoi(id)
	user := repo.GetUser(c.Request().Context(), idStr)
	orders := repo.GetMultiOrders(c.Request().Context(), user.ID)
	c.JSON(200, orders)
	return nil
}
