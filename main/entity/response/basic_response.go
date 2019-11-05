package response

type BasicResponse struct {
	ResponseCode string `json:"responseCode" bson:"responseCode"`
	Message      string `json:"message" bson:"response"`
}