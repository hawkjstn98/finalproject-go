package request

type AddOrUpdatePhoneRequest struct {
	Username string `json:"username" bson:"username"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
}
