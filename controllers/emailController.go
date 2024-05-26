package controllers

import (
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/d-cryptic/crm-golang-backend/models"
    "github.com/d-cryptic/crm-golang-backend/utils"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailRequest struct {
    From    string `json:"from" binding:"required,email"`
    To      string `json:"to" binding:"required,email"`
    Subject string `json:"subject" binding:"required"`
    Body    string `json:"body" binding:"required"`
}

func SendEmailHandler(c *gin.Context) {
    var req EmailRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Failed to bind JSON: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Received email request: %+v\n", req)

    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
    if err != nil {
        log.Printf("Invalid SMTP port: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid SMTP port"})
        return
    }
    smtpUsername := os.Getenv("SMTP_USERNAME")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

    if smtpHost == "" || smtpPort == 0 || smtpUsername == "" || smtpPassword == "" {
        log.Println("SMTP configuration not set")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "SMTP configuration not set"})
        return
    }
    log.Println("SMTP configuration obtained successfully")

    smtpConfig := utils.SMTPConfig{
        Host:     smtpHost,
        Port:     smtpPort,
        Username: smtpUsername,
        Password: smtpPassword,
    }

    trackingID := primitive.NewObjectID().Hex()

    tracking := models.EmailTracking{
        ID:       trackingID,
        From:     req.From,
        To:       req.To,
        Subject:  req.Subject,
        Status:   "sent",
    }
    err = tracking.Save()
    if err != nil {
        log.Printf("Failed to save email tracking info: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save email tracking info"})
        return
    }

    err = utils.SendEmail(smtpConfig, req.From, req.To, req.Subject, req.Body, trackingID)
    if err != nil {
        log.Printf("Failed to send email: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
        return
    }
    log.Println("Email sent successfully")

    c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

func TrackOpenHandler(c *gin.Context) {
    trackingID := c.Param("trackingID")
    log.Printf("Email opened with tracking ID: %s\n", trackingID)

    err := models.UpdateTrackingStatus(trackingID, "opened")
    if err != nil {
        log.Printf("Failed to update tracking status: %v\n", err)
        c.String(http.StatusInternalServerError, "Failed to track email open")
        return
    }

    c.String(http.StatusOK, "Email open tracked")
}