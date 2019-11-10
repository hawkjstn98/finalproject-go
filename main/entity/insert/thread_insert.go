package insert

import "time"

type ThreadInsert struct {
	Timestamp     time.Time `json:"timestamp" bson:"timestamp"`
	Name          string    `json:"name" bson:"name"`
	Category      string    `json:"category" bson:"category"`
	MakerUsername string    `json:"makerUsername" bson:"makerUsername"`
	MakerImage    string    `json:"makerImage" bson:"makerImage"`
	Description   string    `json:"description" bson:"description"`
	CommentCount  int       `json:"commentCount" bson:"commentCount"`
}
