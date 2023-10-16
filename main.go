package main

import (
	"BankAuthenticationProject/api"
	"BankAuthenticationProject/router"
	"context"
)

//var (
//	databaseCollection        *mongo.Collection
//	S3Sess           *session.Session
//	RabbitConnection *amqp.Connection
//)

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

	e.Start(":8000")
}
