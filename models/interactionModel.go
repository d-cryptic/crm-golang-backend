package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Interaction represents a scheduled interaction (meeting) with a customer
type Interaction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CustomerID primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title      string             `json:"title"`
	Details    string             `json:"details"`
	ScheduledAt time.Time         `json:"scheduled_at"`
	CreatedAt  time.Time          `json:"created_at"`
}

// TODO - create separate path for users - also customer and userid should not be same