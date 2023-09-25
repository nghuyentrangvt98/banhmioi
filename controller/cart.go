package controller

import (
	"log/slog"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nghuyentrangvt98/banhmioi/model"
	"github.com/nghuyentrangvt98/banhmioi/repo"
)

func CreateCart(c echo.Context) error {
	var cartModel model.CartCreate
	err := c.Bind(&cartModel)
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
	cartModel.UserID = idStr
	cart := repo.CreateCart(c.Request().Context(), cartModel)
	c.JSON(200, cart)
	return nil
}

func GetCarts(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id, err := claims.GetSubject()
	if err != nil {
		slog.Error(err.Error())
	}
	idStr, _ := strconv.Atoi(id)
	user := repo.GetUser(c.Request().Context(), idStr)
	carts := repo.GetMultiCarts(c.Request().Context(), user.ID)
	c.JSON(200, carts)
	return nil
}
