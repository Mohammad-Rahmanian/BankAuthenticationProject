package utils

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

var (
	RabbitConnection *amqp.Connection
	RabbitChannel    *amqp.Channel
	RabbitQueue      amqp.Queue
	RabbitDelivery   <-chan amqp.Delivery
)

func CreateChannel() error {
	var err error
	RabbitChannel, err = RabbitConnection.Channel()
	if err != nil {
		logrus.Println("Can not create channel")
		return err
	}
	RabbitQueue, err = RabbitChannel.QueueDeclare("user_ids", false, true, false, true, nil)
	if err != nil {
		logrus.Println("Can not create queue")
		return err
	}
	err = RabbitChannel.QueueBind(RabbitQueue.Name, "#", "amq.topic", false, nil)
	if err != nil {
		logrus.Println("Can not bind")
		return err
	}
	RabbitDelivery, err = RabbitChannel.Consume(RabbitQueue.Name, "", false, false, false, true, nil)
	if err != nil {
		logrus.Println("Can not create consume queue")
		return err
	}
	return nil
}
func ConnectMQ() error {
	url := "amqps://zftmeqcg:qfY06WfgC7G94ne1CXvPnPSvOg5ZMJAK@hummingbird.rmq.cloudamqp.com/zftmeqcg"
	var err error
	RabbitConnection, err = amqp.Dial(url)
	if err != nil {
		logrus.Warnln("Can not connect to MQ ", err)
		return err
	}
	logrus.Println("Connected to MQ")
	return err
}

func CloseMQ() {
	err := RabbitConnection.Close()
	if err != nil {
		logrus.Warnln("Can not close MQ ", err)
		logrus.Println(err)
	}
}

func WriteMQ(message string) error {
	channel, _ := RabbitConnection.Channel()
	msg := amqp.Publishing{
		DeliveryMode: 1,
		ContentType:  "text/plain",
		Body:         []byte(message),
	}

	err := channel.PublishWithContext(context.TODO(), "amq.topic", "ping", false, false, msg)
	if err != nil {
		logrus.Warnln("Can not write to Mq ", err)
		return err
	}
	return nil
}
func ReadMQ() (string, error) {
	for message := range RabbitDelivery {
		err := message.Ack(false)
		if err != nil {
			return "", err
		}
		return string(message.Body), nil
	}
	return "", nil
}
