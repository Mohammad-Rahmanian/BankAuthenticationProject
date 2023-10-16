package router

import (
	"BankAuthenticationProject/api/handlers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", handlers.RegisterRequest)
	e.GET("/check", handlers.CheckRequest)

	// create groups
	//adminGroup := e.Group("/admin")

	//e.Use(middleware.Static("./static"))

	//adminGroup.Use(middleware.Logger())
	//adminGroup.Use(middleware.BasicAuth(func(userName string, passWord string, context echo.Context) (bool, error) {
	//	if userName == "Sajjad" && passWord == "1244" {
	//		return true, nil
	//	} else {
	//		return false, nil
	//	}
	//}))

	//api.MainGroup(e)
	//api.AdminGroup(adminGroup)
	return e
}
