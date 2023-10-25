package utils

import (
	"BankAuthenticationProject/configs"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func ConnectMongo() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(configs.DatabaseURL).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		logrus.Errorf("errror is : %v\n", err)
	}
	logrus.Println("Connected to MongoDB")

	collection = client.Database("Authenticator").Collection("users")
	return client
}

func SendPing(client *mongo.Client) {
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		logrus.Println(err)
	}
	logrus.Println("database is fine!")
}

func Insert(user User) error {
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		logrus.Println("User is already registered")
		return err
	}
	return err
}

func UpdateState(encryptedNationalId, state string) error {
	update := bson.D{
		{"$set", bson.D{
			{"state", state},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", encryptedNationalId}}, update)
	if err != nil {
		return err
	}
	logrus.Println("User state is updated")
	return nil
}

func FindUser(encryptedNationalId string) (*User, error) {
	var user User
	err := collection.FindOne(context.TODO(), bson.D{{"_id", encryptedNationalId}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
