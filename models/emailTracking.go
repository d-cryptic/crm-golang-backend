package models

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmailTracking struct {
    ID      string `bson:"_id"`
    From    string `bson:"from"`
    To      string `bson:"to"`
    Subject string `bson:"subject"`
    Status  string `bson:"status"`
}

func (e *EmailTracking) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := getEmailTrackingCollection()
    _, err := collection.InsertOne(ctx, e)
    if err != nil {
        log.Printf("Failed to save email tracking info: %v\n", err)
        return err
    }
    return nil
}

func UpdateTrackingStatus(trackingID, status string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := getEmailTrackingCollection()
    filter := bson.M{"_id": trackingID}
    update := bson.M{"$set": bson.M{"status": status}}
    _, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
    if err != nil {
        log.Printf("Failed to update tracking status: %v\n", err)
        return err
    }
    return nil
}

func getEmailTrackingCollection() *mongo.Collection {
    client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
    if err != nil {
        log.Fatalf("Failed to create MongoDB client: %v", err)
    }

    err = client.Connect(context.Background())
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    return client.Database("crm").Collection("email_tracking")
}
