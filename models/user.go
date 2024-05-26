package models

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type singleUser struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

func GetUserByID(userID primitive.ObjectID) (*singleUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := getUserCollection()
	filter := bson.M{"_id": userID}
	var user singleUser
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to get user by ID: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func getUserCollection() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, _ := mongo.NewClient(clientOptions)
	return client.Database("myDatabase").Collection("users")
}
