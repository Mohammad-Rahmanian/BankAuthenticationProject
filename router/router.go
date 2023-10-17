package router

import (
	"BankAuthenticationProject/api/handlers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", handlers.RegisterRequest)
	e.GET("/check", handlers.CheckRequest)
	//e.Use(middleware.Logger())
	return e
}
