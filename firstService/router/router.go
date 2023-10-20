package router

import (
	"BankAuthenticationProject/firstService/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/register", handlers.RegisterRequest)
	e.GET("/check", handlers.CheckRequest)
	return e
}
