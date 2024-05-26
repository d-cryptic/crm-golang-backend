package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/d-cryptic/crm-golang-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNotification handles creating a new notification
func CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	notification.ID = primitive.NewObjectID()
	notification.CreatedAt = time.Now()

	if err := notification.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		log.Printf("%v", err) // Convert err to a string before passing it to log.Printf()
		return
	}

	c.JSON(http.StatusOK, notification)
}

// GetNotifications handles fetching notifications
func GetNotifications(c *gin.Context) {
	notifications, err := models.GetUpcomingNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}
