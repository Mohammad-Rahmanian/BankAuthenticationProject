package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Email      string `json:"email"`
	Lastname   string `json:"lastname"`
	NationalID int    `json:"nationalID"`
	IP         string `json:"IP"`
	Image1     string `json:"image1"`
	Image2     string `json:"image2"`
	State      string `json:"state"`
}

func registerRequest(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading body")
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(body, &user)
	log.Printf("%#v", user)
	if err != nil {
		log.Printf("Failed unmarshlling body" + string(err.Error()))
		return c.String(http.StatusInternalServerError, "")
	}
	return c.String(http.StatusInternalServerError, "")
	//log.Printf("%#v", cat)
	//return c.String(http.StatusOK, "Your authentication request has been registered.")
}
func checkRequest(c echo.Context) {
	//nationalID := c.QueryParam("nationalID")
	//if dataType == "string" {
	//	return c.String(http.StatusOK, fmt.Sprintf("Your cat name is "+string(catName)+"  Your cat type is "+catType))
	//}
	//if dataType == "json" {
	//	return c.JSON(http.StatusOK, map[string]string{
	//		"name": catName,
	//		"type": catType,
	//	})
	//}
	//return c.JSON(http.StatusBadRequest, map[string]string{
	//	"error": " your need to let us want json or string",
	//})
}
func main() {
	e := echo.New()

	e.POST("/register", registerRequest)

	e.Start(":8000")

}
