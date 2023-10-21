package router

import (
	"BankAuthenticationProject/firstService/api/handlers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/register", handlers.RegisterRequest)
	e.GET("/check", handlers.CheckRequest)
	return e
}
