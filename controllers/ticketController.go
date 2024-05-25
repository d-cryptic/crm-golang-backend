package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/d-cryptic/crm-golang-backend/config"
	"github.com/d-cryptic/crm-golang-backend/models"
)

var ticketCollection *mongo.Collection = config.GetCollection(config.ConnectDB(), "tickets")

func CreateTicket(c *gin.Context) {
	log.Println("Creating ticket...")
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.ID = primitive.NewObjectID()
	ticket.Status = "open"
	ticket.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := ticketCollection.InsertOne(ctx, ticket)
	if err != nil {
		log.Println("Error creating ticket in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating ticket"})
		return
	}

	log.Println("Ticket created successfully")
	c.JSON(http.StatusCreated, ticket)
}

func ResolveTicket(c *gin.Context) {
	log.Println("Resolving ticket...")
	ticketID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(ticketID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"status":      "resolved",
			"resolved_at": time.Now(),
		},
	}

	_, err := ticketCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Println("Error resolving ticket:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while resolving ticket"})
		return
	}

	log.Println("Ticket resolved successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Ticket resolved successfully"})
}
