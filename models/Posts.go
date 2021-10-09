package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Caption  string             `json:"caption" bson:"caption"`
	ImageUrl string             `json:"imageUrl" bson:"imageUrl"`
	Author   string             `json:"author" bson:"author"`
	Time     time.Time          `json:"time" bson:"time"`
}
