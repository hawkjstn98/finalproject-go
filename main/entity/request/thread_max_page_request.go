package request

type ThreadMaxPageRequest struct {
	Category string `json:"category" bson:"category"`
	Page int `json:"page" bson:"page"`
}
