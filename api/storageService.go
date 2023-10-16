package api

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"mime/multipart"
)

var StorageSession *session.Session
var RabbitConnection *amqp.Connection

func ConnectS3() *session.Session {
	var err error
	StorageSession, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("175b85c0-26e0-4ca2-b25a-e2d3ec0b56cb", "380b008e0005cbf7f14e68c6e365f65940116d3016b182297f73bca9d8f17e4f", ""),
		Region:      aws.String("default"),
		Endpoint:    aws.String("https://sajjadstorage.s3.ir-thr-at1.arvanstorage.ir"),
	})

	if err != nil {
		logrus.Warnln("can not connect to s3", err)
	}
	logrus.Infoln("connected to S3 instance")

	return StorageSession
}

func UploadS3(sess *session.Session, fileHeader *multipart.FileHeader, bucket string, ID string) string {
	uploader := s3manager.NewUploader(sess)
	file, err := fileHeader.Open()
	key := fmt.Sprintf("%s", fileHeader.Filename+ID)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		logrus.Warnln("Unable to upload %q to %q, %v", fileHeader.Filename, bucket, err)
	}
	logrus.Infoln("Successfully uploaded %q to %q\n", fileHeader.Filename, bucket)

	return key
}
func ConnectMQ() {
	url := "amqps://kkyazhsy:tXIS6A9botHvCsdygKjRfM3FwFR4qElg@hummingbird.rmq.cloudamqp.com/kkyazhsy"
	var err error
	RabbitConnection, err = amqp.Dial(url)
	if RabbitConnection == nil {
		println("******************")
	}
	if err != nil {
		logrus.Println(err)
	}
}

func CloseMQ(connection *amqp.Connection) {
	err := connection.Close()
	if err != nil {
		return
	}
}

func WriteMQ(message string) error {
	//if RabbitConnection == nil {
	//	println("******************")
	//}
	channel, _ := RabbitConnection.Channel()
	msg := amqp.Publishing{
		DeliveryMode: 1,
		ContentType:  "text/plain",
		Body:         []byte(message),
	}

	err := channel.PublishWithContext(context.TODO(), "amq.topic", "ping", false, false, msg)
	if err != nil {
		logrus.Warnln("cant publish message to queue")
		return err
	}

	return nil
}
