package response

type ThreadMaxPageResponse struct {
	Response BasicResponse `json:"response" bson:"response"`
	Page int `json:"page" bson:"page"`
}
