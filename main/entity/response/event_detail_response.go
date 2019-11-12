package response

import "github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"

type EventDetailResponse struct {
	Response BasicResponse `json:"response" bson:"response"`
	Event    *event.GameEvent `json:"event" bson:"event"`
	Distance float32 `json:"distance" bson:"distance"`
}
