package router

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/nghuyentrangvt98/banhmioi/config"
	"github.com/nghuyentrangvt98/banhmioi/controller"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Routing(e *echo.Echo) {
	e.GET("/", hello, echojwt.JWT(config.SECRETKEY))
	e.POST("/users", controller.CreateUser)
	e.POST("/login", controller.Login)
	e.GET("/products", controller.GetMultiProduct)
	e.GET("/products/:id", controller.GetProduct)
	e.GET("/news", controller.GetMultiNews)
	e.GET("/news/:id", controller.GetNews)
}
