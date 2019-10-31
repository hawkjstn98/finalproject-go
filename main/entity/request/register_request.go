package request

type RegisterRequest struct {
	PhoneNumber 	string  `json:"phoneNumber" bson:"phoneNumber"`
	Email 			string `json:"email" bson:"email"`
	Username 		string `json:"username" bson:"username"`
	Password        string `json:"password" bson:"password"`
}
