package response

import "github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"

type EventHomeResponse struct {
	Response BasicResponse
	EventList []*event.GameEvent `json:"eventList" bson:"eventList"`
}
