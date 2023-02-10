package database

import (
	"context"
	"fmt"
	"log"

	"github.com/FadyGamilM/gomongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOne(movie models.Movie) {
	insertionResult, err := Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("One item is inserted in DB with id => ", insertionResult.InsertedID)
}

func UpdateOneById(id string) {
	// convert string id into Objectid
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Error while parsing the given id, [ERROR] => ", err)
	}

	filter := bson.M{"_id": objId}

	// how i will update this filtered row
	update := bson.M{"$set": bson.M{"watched": true}}

	updationResult, UpdationErr := Collection.UpdateOne(context.Background(), filter, update)
	if UpdationErr != nil {
		log.Fatal("Error while updating the item, [ERROR] => ", UpdationErr)
	}
	fmt.Println("Number of modified entities => ", updationResult.ModifiedCount)
}

func DeleteById(id string) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Error while parsing the given id, [ERROR] => ", err)
	}
	filter := bson.M{"_id": objId}
	DeletionResult, DeletionError := Collection.DeleteOne(context.Background(), filter)

	if DeletionError != nil {
		log.Fatal("Error while Deleting the item, [ERROR] => ", DeletionError)
	}
	fmt.Println("Number of deleted entities => ", DeletionResult.DeletedCount)
}

func DeleteAll() {
	DeletionResult, DeletionError := Collection.DeleteMany(context.Background(), bson.M{}, nil)
	if DeletionError != nil {
		log.Fatal("Error while Deleting the item, [ERROR] => ", DeletionError)
	}
	fmt.Println("Number of deleted entities => ", DeletionResult.DeletedCount)
}

func GetAll() []primitive.M {
	cursor, err := Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

// func GetById(id string) models.Movie {
// 	// convert the id
// 	objId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// look for this entity and return it
// 	filter := bson.M{"_id": objId}
// 	movie := Collection.FindOne(context.Background(), filter)
// 	return movie
// }
