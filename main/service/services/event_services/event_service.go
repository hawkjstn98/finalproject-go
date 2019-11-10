package event_services

import (
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/event_repository"
)

func GetHome(req *request.EventHomeRequest) (res string, err error){
	if req.Page < 1 {
		return "", fmt.Errorf("invalid paging")
	}
	gameEvent, err := event_repository.GetEventHome(req.Page)
	if err != nil {
		return
	}
	var resp response.EventHomeResponse
	eventList := MapToEventList(gameEvent)
	resp.EventList = eventList
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	b, err := json.Marshal(resp)
	return string(b), nil
}

func MapToEventList(events []*event.GameEvent) (eventList []*event.GameEvent){
	for _, e := range events{
		e.Timestamp = e.ID.Timestamp()
		eventList = append(eventList, e)
	}
	return
}