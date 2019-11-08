package request

type ThreadDetailRequest struct {
	ThreadID string `json:"threadId" bson:"threadId"`
	Page int `json:"page" bson:"page"`
}
