package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the structure of a user in the CRM system
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"-"`
	Company  string             `json:"company,omitempty"`
	Status   string             `json:"status"`
	Notes    string             `json:"notes,omitempty"`
}