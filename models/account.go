package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Account struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Login    string        `bson:"login"`
	Password string        `bson:"password"`
}
