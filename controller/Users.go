package controller

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LegalId    string             `json:"legalId"`
	Name       string             `json:"name"`
	Profession string             `json:"profession"`
	Gender     string             `json:"gender"`
}
