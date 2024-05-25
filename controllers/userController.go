package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/d-cryptic/crm-golang-backend/config"
	"github.com/d-cryptic/crm-golang-backend/models"
	"github.com/d-cryptic/crm-golang-backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.ConnectDB(), "users")

func CreateUser(c *gin.Context) {
	log.Println("Creating user...")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		log.Println("User with email", user.Email, "already exists")
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		log.Println("Error checking existing user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while checking existing user"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}
	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()
	user.Status = "active"

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Error creating user in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	log.Println("User created successfully")
	c.JSON(http.StatusCreated, user)
}

