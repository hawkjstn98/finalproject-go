package forum

import "go.mongodb.org/mongo-driver/bson/primitive"

type ObjectComment struct {
	Timestamp  	 	primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Username 		string `json:"username" bson:"username"`
	ThreadComment	string `json:"threadComment" bson:"threadComment"`
}

