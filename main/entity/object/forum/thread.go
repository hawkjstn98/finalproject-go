package forum

import "go.mongodb.org/mongo-driver/bson/primitive"

const ThreadCollection = "thread"

type Thread struct {
	Id				primitive.ObjectID `json:"id" bson:"_id"`
	Timestamp  	 	primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Name			string `json:"name" bson:"name"`
	Category		string `json:"category" bson:"category"`
	MakerUsername	string `json:"makerUsername" bson:"makerUsername"`
	MakerImage		string `json:"makerImage" bson:"makerImage"`
	Description		string `json:"description" bson:"description"`
	CommentCount	int `json:"commentCount" bson:"commentCount"`
}

