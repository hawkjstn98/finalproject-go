package request

import (
	"time"
)

type CreateThreadRequest struct {
	Timestamp     time.Time `json:"timestamp" bson:"timestamp"`
	Name          string    `json:"name" bson:"name"`
	Category      string    `json:"category" bson:"category"`
	MakerUsername string    `json:"makerUsername" bson:"makerUsername"`
	Description   string    `json:"description" bson:"description"`
}

type CreateThreadCommentRequest struct {
	MasterThreadID string    `json:"threadMasterId" bson:"threadMasterId"`
	Timestamp      time.Time `json:"timestamp" bson:"timestamp"`
	Category       string    `json:"category" bson:"category"`
	MakerUsername  string    `json:"makerUsername" bson:"makerUsername"`
	MakerImage     string    `json:"makerImage" bson:"makerImage"`
	Description    string    `json:"description" bson:"description"`
}
