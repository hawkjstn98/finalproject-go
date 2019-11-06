package request

type ThreadCategoryRequest struct {
	Category string `json:"category" bson:"category"`
}