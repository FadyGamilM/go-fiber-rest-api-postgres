package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connString     string = `mongodb+srv://fady:fadyfady@sandbox.oxpp9.mongodb.net/?retryWrites=true&w=majority`
	dbName                = "moviesDB"
	collectionName        = "watchingList"

// reference to the collection
)

var Collection *mongo.Collection

// connect to the database
// run at first time the app is booted up and only one time
func Init() {
	// client option conn
	clientOption := options.Client().ApplyURI(connString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to mongodb is done")

	// set the reference
	Collection = (*mongo.Collection)(client.Database(dbName).Collection(collectionName))

	fmt.Println("collection instance is ready")
}
