package request

type AddOrUpdateProfileImage struct {
	ImageInString string `json:"imageInString" bson:"imageInString"`
	Username string `json:"username" bson:"username"`
}
