package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Session struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Token     string        `bson:"token"`
	AccountID bson.ObjectID `bson:"account_id"`
	ExpiresAt time.Time     `bson:"expires_at"`
}
