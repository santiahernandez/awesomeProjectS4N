package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LegalId    string             `json:"curso"`
	Name       string             `json:"salon"`
	Profession string             `json:"profesor"`
	Gender     string             `json:"horario"`
}