package main

import (
	"BankAuthenticationProject/secondService/api"
	"BankAuthenticationProject/utils"
	"context"
	"github.com/sirupsen/logrus"
)

func main() {
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
	err = utils.CreateChannel()
	if err != nil {
		logrus.Println(err)
	}

	for {
		api.Authenticate()
	}

}
