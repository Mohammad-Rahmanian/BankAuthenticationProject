package api

import (
	"BankAuthenticationProject/utils"
	"github.com/sirupsen/logrus"
)

func Authenticate() {
	println("Read from MQ .......")
	encryptedNationalId, err := utils.ReadMQ()
	if err != nil {
		logrus.Println("Can not read from MQ")
	} else {
		println("Complete.")
	}
	user, err := utils.FindUser(encryptedNationalId)
	if err != nil {
		logrus.Println("Can not find user")

	}
	firstImage, err := utils.DownloadS3(utils.StorageSession, "SajjadStorage", user.FirstImage)
	if err != nil {
		logrus.Println("Can not download first image:", err)
	}
	SecondImage, err := utils.DownloadS3(utils.StorageSession, "SajjadStorage", user.SecondImage)
	if err != nil {
		logrus.Println("Can not download second image:", err)
	}

	firstID, err := FaceDetection(firstImage)
	secondID, err := FaceDetection(SecondImage)
	similarityScore, err := FaceSimilarity(firstID, secondID)
	if err != nil {
		logrus.Println(err)
	}
	logrus.Printf("Face similarity score: %d\n", similarityScore)
	if similarityScore >= 80 {
		err := utils.UpdateState(encryptedNationalId, "accepted")
		if err != nil {
			logrus.Println("Can not update user state", err)
		}
		user.State = "accepted"

		_, err = SendMail(user.Email, user.State)
		if err != nil {
			logrus.Println("Can not send mail:", err)
		} else {
			logrus.Println("Mail has been sent successfully ")
		}
		logrus.Printf("%s has been authenticated successfully.", user.Lastname)

	} else {
		err := utils.UpdateState(encryptedNationalId, "rejected")
		if err != nil {
			logrus.Println("Can not update user state")
		}
		user.State = "rejected"
		_, err = SendMail(user.Email, user.State)
		if err != nil {
			logrus.Println("Can not send Email:", err)
		} else {
			logrus.Println("Mail has been sent successfully ")
		}
		logrus.Printf("%s has not been authenticated successfully. Please try again", user.Lastname)

	}

}
