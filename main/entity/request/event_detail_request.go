package request

type EventDetailRequest struct {
	EventId string `json:"eventId" bson:"eventId"`
}
