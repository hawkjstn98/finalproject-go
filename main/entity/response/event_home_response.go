package response

import "github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"

type EventHomeResponse struct {
	Response BasicResponse `json:"response" bson:"response"`
	MaxPage int64 `json:"maxPage" bson:"maxPage"`
	EventList []*event.GameEvent `json:"eventList" bson:"eventList"`
}
