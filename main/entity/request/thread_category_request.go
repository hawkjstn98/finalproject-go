package request

type ThreadCategoryRequest struct {
	Category string `json:"category" bson:"category"`
	Page int `json:"page" bson:"page"`
}