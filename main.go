package main

import (
	"BankAuthenticationProject/api"
	"BankAuthenticationProject/router"
	"context"
)

func main() {
	e := router.New()
	client := api.ConnectMongo()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	api.ConnectS3()
	api.ConnectMQ()
	defer api.CloseMQ()
	e.Start(":8000")
}
