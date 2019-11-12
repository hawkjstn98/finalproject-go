package request

type EventDetailRequest struct {
	EventId       string `json:"eventId" bson:"eventId"`
	UserLatitude  string `json:"userLatitude" bson:"userLatitude"`
	UserLongitude string `json:"userLongitude" bson:"userLongitude"`
}
