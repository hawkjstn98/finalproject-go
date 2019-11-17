package request

type MyEventRequest struct {
	Username  string `json:"username" bson:"username"`
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}
