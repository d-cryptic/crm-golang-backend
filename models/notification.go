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

type Notification struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty"`
	Message     string             `bson:"message"`
	CreatedAt   time.Time          `bson:"created_at"`
	NotifiedAt  time.Time          `bson:"notified_at,omitempty"`
	IsNotified  bool               `bson:"is_notified"`
	EventTime   time.Time          `bson:"event_time"`
}

func (n *Notification) Save() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := getNotificationCollection()
	_, err := collection.InsertOne(ctx, n)
	if err != nil {
		log.Printf("Failed to save notification: %v\n", err)
		return err
	}
	return nil
}

func GetUpcomingNotifications() ([]Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := getNotificationCollection()
	filter := bson.M{"is_notified": false, "event_time": bson.M{"$lte": time.Now().Add(1 * time.Hour)}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Failed to get upcoming notifications: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []Notification
	if err := cursor.All(ctx, &notifications); err != nil {
		log.Printf("Failed to decode notifications: %v\n", err)
		return nil, err
	}
	return notifications, nil
}

func MarkAsNotified(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := getNotificationCollection()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_notified": true, "notified_at": time.Now()}}
	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		log.Printf("Failed to mark notification as notified: %v\n", err)
		return err
	}
	return nil
}

func getNotificationCollection() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, _ := mongo.NewClient(clientOptions)
	return client.Database("myDatabase").Collection("notifications")
}
