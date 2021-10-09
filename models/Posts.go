package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	Id       primitive.ObjectID `json:"id" bson:"_id" validate:"nil=false"`
	Caption  string             `json:"caption" bson:"caption"`
	ImageUrl string             `json:"imageUrl" bson:"imageUrl" validate:"nil=false"`
	Author   string             `json:"author" bson:"author" validate:"nil:false"`
	Time     time.Time          `json:"time" bson:"time"`
}
