package handlers

import (
	"BankAuthenticationProject/api"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RegisterRequest(c echo.Context) error {
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	nationalId := c.FormValue("nationalID")
	ip := c.RealIP()
	firstImage, err := c.FormFile("firstImage")
	if err != nil {
		logrus.Printf("Unable to open first iamge")
		return c.JSON(http.StatusBadRequest, "Unable to open file")
	}

	secondImage, err := c.FormFile("secondImage")
	if err != nil {
		logrus.Printf("Unable to open second image")
		return c.JSON(http.StatusBadRequest, "Unable to open file")
	}

	nationalId = base64.StdEncoding.EncodeToString([]byte(nationalId))
	firstPath := api.UploadS3(api.StorageSession, firstImage, "SajjadStorage", lastname)
	secondPath := api.UploadS3(api.StorageSession, secondImage, "SajjadStorage", lastname)
	user := api.NewUSer(email, lastname, nationalId, ip, firstPath, secondPath, "pending")
	err = api.Insert(email, lastname, nationalId, ip, firstPath, secondPath)
	if err != nil {
		logrus.Printf("Can not insert to database")
		return c.JSON(http.StatusInternalServerError, "Can not insert to database")
	}
	err = api.WriteMQ(nationalId)
	if err != nil {
		logrus.Printf("Can not write to queue")
		return c.JSON(http.StatusInternalServerError, "Can not write to queue")
	}
	fmt.Printf("%#v", user)
	return c.String(http.StatusOK, "Your authentication request has been registered.")
}

func CheckRequest(c echo.Context) error {
	nationalID := c.QueryParam("nationalID")
	nationalID = base64.StdEncoding.EncodeToString([]byte(nationalID))
	api.GetAll()
	var user = api.FindUser(nationalID)

	if user.IP != c.RealIP() {
		return c.String(403, "Your access is unauthorized")
	}
	if user.State == "pending" {
		return c.String(http.StatusOK, fmt.Sprintf("Pending...."))

	} else if user.State == "" {
		return c.String(http.StatusOK, fmt.Sprintf("Pending...."))

	} else if user.State == "rejected" {
		return c.String(http.StatusOK, fmt.Sprintf("Your authentication request has been rejected. Please try again later."))

	} else if user.State == "accepted" {
		return c.String(http.StatusOK, fmt.Sprintf("Your authentication request has been accepted. Your username is "+string(user.Lastname)))
	} else {
		return c.String(http.StatusBadRequest, fmt.Sprintf("heb"))
	}

}
