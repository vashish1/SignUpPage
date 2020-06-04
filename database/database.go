package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Client) {

	url:=os.Getenv("MongoUrl")

	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	usercollection := client.Database("Signup").Collection("User")
	return usercollection, client
}

//Insertintouserdb inserts the data into the database
func Insertintouserdb(usercollection *mongo.Collection, u User) {

	// fmt.Println(u.Username)
	insertResult, err := usercollection.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Println("THe error is", err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

//Findfromuserdb finds the required data
func Findfromuserdb(usercollection *mongo.Collection, st string) bool {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//Finddb finds the required database
func Finddb(c *mongo.Collection, s string) User {
	filter := bson.D{primitive.E{Key: "email", Value: s}}
	var result User

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Print(err)
		return result
	}
	return result
}
