package request

type EventHomeRequest struct {
	Page      int    `json:"page" bson:"page"`
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}
