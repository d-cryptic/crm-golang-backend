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

var interactionCollection *mongo.Collection = config.GetCollection(config.ConnectDB(), "interactions")

func ScheduleInteraction(c *gin.Context) {
	log.Println("Scheduling interaction...")
	var interaction models.Interaction
	if err := c.ShouldBindJSON(&interaction); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interaction.ID = primitive.NewObjectID()
	interaction.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := interactionCollection.InsertOne(ctx, interaction)
	if err != nil {
		log.Println("Error scheduling interaction in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while scheduling interaction"})
		return
	}

	log.Println("Interaction scheduled successfully")
	c.JSON(http.StatusCreated, interaction)
}

func GetInteractionsByCustomerID(c *gin.Context) {
	log.Println("Getting interactions for customer...")
	customerID := c.Param("customer_id")
	objID, _ := primitive.ObjectIDFromHex(customerID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var interactions []models.Interaction
	cursor, err := interactionCollection.Find(ctx, bson.M{"customer_id": objID})
	if err != nil {
		log.Println("Error finding interactions:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while finding interactions"})
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &interactions); err != nil {
		log.Println("Error decoding interactions:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding interactions"})
		return
	}

	log.Println("Interactions retrieved successfully")
	c.JSON(http.StatusOK, interactions)
}
