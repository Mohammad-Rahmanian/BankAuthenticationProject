package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
)

var StorageSession *session.Session

func ConnectS3() error {
	var err error
	StorageSession, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("175b85c0-26e0-4ca2-b25a-e2d3ec0b56cb", "380b008e0005cbf7f14e68c6e365f65940116d3016b182297f73bca9d8f17e4f", ""),
		Region:      aws.String("default"),
		Endpoint:    aws.String("https://sajjadstorage.s3.ir-thr-at1.arvanstorage.ir"),
	})

	if err != nil {
		logrus.Warnln("Can not connect to s3 ", err)
		return err
	}
	logrus.Println("connected to S3")

	return err
}

func UploadS3(sess *session.Session, fileHeader *multipart.FileHeader, bucket string, ID string) (string, error) {
	uploader := s3manager.NewUploader(sess)
	file, err := fileHeader.Open()
	key := fmt.Sprintf("%s", fileHeader.Filename+ID)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		logrus.Printf("Unable to upload %v to %v, %v\n", fileHeader.Filename, bucket, err)
		return "", err
	}
	logrus.Printf("Successfully uploaded %v to %v\n", fileHeader.Filename, bucket)

	return key, err
}
func DownloadS3(sess *session.Session, bucket string, key string) (*os.File, error) {
	dir, _ := os.Getwd()
	file, err := os.Create(dir + "/images/" + key)
	if err != nil {
		logrus.Println("Can not open file:", err)
		return nil, err
	}
	s3Client := s3.New(sess)

	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	_, err = io.Copy(file, obj.Body)
	if err != nil {
		logrus.Println("Can not copy file:", err)
		return nil, err
	}

	return file, nil
}
