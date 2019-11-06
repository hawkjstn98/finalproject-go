package request

type ThreadDetailRequest struct {
	ThreadID int `json:"threadId" bson:"threadId"`
	Page int `json:"page" bson:"page"`
}
