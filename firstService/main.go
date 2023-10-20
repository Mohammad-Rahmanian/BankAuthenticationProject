package main

import (
	"BankAuthenticationProject/firstService/router"
	"BankAuthenticationProject/utils"
	"context"
	"github.com/sirupsen/logrus"
)

func main() {
	e := router.New()
	client := utils.ConnectMongo()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	err := utils.ConnectS3()
	if err != nil {
		logrus.Println(err)
	}
	err = utils.ConnectMQ()
	if err != nil {
		logrus.Println(err)
	}
	defer utils.CloseMQ()
	err = e.Start(":8016")
	if err != nil {
		logrus.Println(err)
	}
}
