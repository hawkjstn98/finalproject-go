package response

import "github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"

type MyEventResponse struct {
	Response  BasicResponse
	EventList []event.GameEvent
}
