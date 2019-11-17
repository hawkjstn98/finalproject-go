package request

type ThreadDetailRequest struct {
	ThreadID string `json:"threadId" bson:"threadId"`
	Page     int64    `json:"page" bson:"page"`
}
