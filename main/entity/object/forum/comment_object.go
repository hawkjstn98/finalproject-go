package forum

import (
	"time"
)

const CommentCollection = "comment"

type ObjectComment struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	ThreadMasterID string             `json:"masterThreadId" bson:"masterThreadId"`
	Timestamp      time.Time          `json:"timestamp" bson:"timestamp"`
	Username       string             `json:"username" bson:"username"`
	ProfileImage   string             `json:"profileImage" bson:"profileImage"`
	ThreadComment  string             `json:"threadComment" bson:"threadComment"`
}
