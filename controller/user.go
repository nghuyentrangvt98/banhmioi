package controller

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nghuyentrangvt98/banhmioi/config"
	"github.com/nghuyentrangvt98/banhmioi/model"
	"github.com/nghuyentrangvt98/banhmioi/repo"
)

func CreateUser(c echo.Context) error {
	var userModel model.UserCreate
	err := c.Bind(&userModel)
	if err != nil {
		c.JSON(400, model.Exception{Message: err.Error()})
		return nil
	}
	user := repo.CreateUser(c.Request().Context(), userModel)
	c.JSON(200, user)
	return nil
}

func Login(c echo.Context) error {
	var userModel model.UserLogin
	err := c.Bind(&userModel)
	if err != nil {
		return err
	}
	user := repo.GetUserByUserName(c.Request().Context(), userModel.Username)
	if user == nil {
		c.JSON(400, model.Exception{Message: "user not exists"})
		return nil
	}
	if user.Password != userModel.Password {
		c.JSON(400, model.Exception{Message: "wrong password"})
		return nil
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "username"
	jwtToken, err := token.SignedString(config.SECRETKEY)
	if err != nil {
		return err
	}

	c.JSON(200, model.FromSchemaUser(user, jwtToken))
	return nil
}
