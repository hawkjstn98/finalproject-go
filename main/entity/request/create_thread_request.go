package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateThreadRequest struct {
	Timestamp  	 	primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Name			string `json:"name" bson:"name"`
	Category		string `json:"category" bson:"category"`
	MakerUsername	string `json:"makerUsername" bson:"makerUsername"`
	Description		string `json:"description" bson:"description"`
}