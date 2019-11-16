package bookmark

type ObjectBookmark struct {
	UserID string `json:"userId" bson:"userId"`
	EventID string `json:"eventId" bson:"eventId"`
}
