package request

import (
	"time"
)

type CreateThreadRequest struct {
	Timestamp  	 	time.Time `json:"timestamp" bson:"timestamp"`
	Name			string `json:"name" bson:"name"`
	Category		string `json:"category" bson:"category"`
	MakerUsername	string `json:"makerUsername" bson:"makerUsername"`
	Description		string `json:"description" bson:"description"`
}