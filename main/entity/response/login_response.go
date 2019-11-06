package response

type LoginResponse struct{
	Response BasicResponse
	Username string `json:"username" bson:"username"`
}