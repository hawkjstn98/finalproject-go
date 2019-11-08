package forum

import "go.mongodb.org/mongo-driver/bson/primitive"

const CommentCollection = "comment"

type ObjectComment struct {
	ThreadMasterID 	string `json:"threadMasterId" bson:"threadMasterId"`
	Timestamp  	 	primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Username 		string `json:"username" bson:"username"`
	ProfileImage	string `json:"profileImage" bson:"profileImage"`
	ThreadComment	string `json:"threadComment" bson:"threadComment"`
}

