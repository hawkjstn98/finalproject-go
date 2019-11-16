package response

type LoginResponse struct{
	Response BasicResponse `json:"response" bson:"response"`
	Username string `json:"username" bson:"username"`
}