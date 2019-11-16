package request

const BookmarkCollection = "bookmark"

type BookmarkRequest struct {
	UserID string `json:"userId" bson:"userId"`
	EventID string `json:"eventId" bson:"eventId"`
}
