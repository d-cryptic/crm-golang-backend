package utils

import (
	"net/smtp"
	"os"

	"github.com/d-cryptic/crm-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SendNotification(userID primitive.ObjectID, message string) error {
	user, err := models.GetUserByID(userID)
	if err != nil {
		return err
	}

	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	to := user.Email
	subject := "Upcoming Event Notification"

	body := "Subject: " + subject + "\n\n" + message

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
