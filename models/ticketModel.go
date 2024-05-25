package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Ticket represents a support ticket raised by a customer
type Ticket struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CustomerID primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	Title      string             `json:"title"`
	Description string            `json:"description"`
	Status     string             `json:"status"` // "open" or "resolved"
	CreatedAt  time.Time          `json:"created_at"`
	ResolvedAt time.Time          `json:"resolved_at,omitempty"`
}
