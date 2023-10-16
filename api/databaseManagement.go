package api

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection

func ConnectMongo() *mongo.Client {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://Mohammad5070:Sajjad5070@hw1cloud.yttxe75.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection = client.Database("validator").Collection("users")
	return client
}

func SendPing(client *mongo.Client) {
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		logrus.Println(err)
	}
	logrus.Println("database is fine!")
}

func Insert(email string, lastname string, nationalId string, ip string, firstImage string, secondImage string) bool {
	err := collection.FindOne(context.TODO(), bson.D{{"nationalId", nationalId}}).Err()
	if err != mongo.ErrNoDocuments {
		logrus.Println("user is already registered: ", nationalId)
		return false
	}
	//nationalId = base64.StdEncoding.EncodeToString([]byte(nationalId))
	user := NewUSer(email, lastname, nationalId, ip, firstImage, secondImage, "pending")
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		logrus.Panicln(err)
	}

	logrus.Println("user is created: ", nationalId, " name:", user.Lastname)

	return true
}

//func Update(nationalId, state string) bool {
//	update := bson.D{
//		{"$set", bson.D{
//			{"state", state},
//		}},
//	}
//	_, err := collection.UpdateOne(context.TODO(), bson.D{{"nationalId", nationalId}}, update)
//	if err != nil {
//		logrus.Warnln("cant update users object")
//		return false
//	}
//
//	return true
//}

func GetAll() []User {
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logrus.Warnln("cant find all users")
	}

	res := make([]User, 0)
	var doc User
	for cur.Next(context.TODO()) {
		err := cur.Decode(&doc)
		if err != nil {
			log.Panicln(err)
		}
		res = append(res, doc)
	}

	return res
}

func FindUser(nationalId string) *User {
	var user User
	err := collection.FindOne(context.TODO(), bson.D{{"nationalID", nationalId}}).Decode(&user)
	if err != nil {
		logrus.Warnln("user not found", err)
	}
	return &user
}
