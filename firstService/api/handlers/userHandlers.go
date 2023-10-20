package handlers

import (
	"BankAuthenticationProject/utils"
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
		logrus.Printf("Unable to open first iamge\n")
		return c.String(http.StatusBadRequest, "Unable to open file")
	}
	secondImage, err := c.FormFile("secondImage")
	if err != nil {
		logrus.Printf("Unable to open second image\n")
		return c.String(http.StatusBadRequest, "Unable to open file")
	}
	encryptedNationalID := base64.StdEncoding.EncodeToString([]byte(nationalId))
	firstPath, err := utils.UploadS3(utils.StorageSession, firstImage, "SajjadStorage", encryptedNationalID)
	if err != nil {
		logrus.Printf("Unable to upload first image\n")
		return c.String(http.StatusBadRequest, "Unable to upload first image")
	}
	secondPath, err := utils.UploadS3(utils.StorageSession, secondImage, "SajjadStorage", encryptedNationalID)
	if err != nil {
		logrus.Printf("Unable to upload first image\n")
		return c.String(http.StatusBadRequest, "Unable to upload first image")
	}
	existUser, _ := utils.FindUser(encryptedNationalID)
	if existUser == nil {
		user := utils.NewUSer(email, lastname, encryptedNationalID, ip, firstPath, secondPath, "pending")
		err = utils.Insert(*user)
		if err != nil {
			logrus.Printf("Can not insert to database\n")
			return c.String(http.StatusInternalServerError, "Can not insert to database")
		}
		err = utils.WriteMQ(encryptedNationalID)
		if err != nil {
			logrus.Printf("Can not write to queue\n")
			return c.String(http.StatusInternalServerError, "Can not write to queue")
		}
		fmt.Printf("%#v\n", *user)
		return c.String(http.StatusOK, "Your authentication request has been registered.")
	} else {
		if existUser.State == "accepted" {
			return c.String(http.StatusOK, "Your authentication request has already been accepted.")
		} else if existUser.State == "Pending" {
			return c.String(http.StatusOK, "Your authentication request has been registered. Please wait to see the result.")
		} else {
			return c.String(http.StatusOK, "Your authentication request has already been rejected")
		}
	}

}

func CheckRequest(c echo.Context) error {
	nationalID := c.QueryParam("nationalID")
	encryptedNationalID := base64.StdEncoding.EncodeToString([]byte(nationalID))
	user, err := utils.FindUser(encryptedNationalID)
	if err != nil {
		return c.String(http.StatusForbidden, "Your national id is wrong")
	} else if user.IP != c.RealIP() {
		return c.String(http.StatusForbidden, "Your access is unauthorized")
	} else if user.State == "pending" {
		return c.String(http.StatusOK, fmt.Sprintf("Pending...."))

	} else if user.State == "rejected" {
		return c.String(http.StatusOK, fmt.Sprintf("Your authentication request has been rejected. Please try again later."))

	} else if user.State == "accepted" {
		return c.String(http.StatusOK, fmt.Sprintf("Your authentication request has been accepted. Your username is "+string(user.Lastname)))
	} else {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Bad request."))
	}

}
