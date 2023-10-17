package main

import (
	"BankAuthenticationProject/api"
	"BankAuthenticationProject/router"
	"context"
	"github.com/sirupsen/logrus"
)

func main() {
	e := router.New()
	client := api.ConnectMongo()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	err := api.ConnectS3()
	if err != nil {
		logrus.Println(err)
	}
	err = api.ConnectMQ()
	if err != nil {
		logrus.Println(err)
	}
	defer api.CloseMQ()
	err = e.Start(":8000")
	if err != nil {
		logrus.Println(err)
	}
}
