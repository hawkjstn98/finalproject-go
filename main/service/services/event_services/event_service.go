package event_services

import (
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/event_repository"
	"log"
)

func GetHome(req *request.EventHomeRequest) (res string, err error){
	if req.Page < 1 {
		return "", fmt.Errorf("invalid paging")
	}
	gameEvent, err := event_repository.GetEventHome(req.Page)
	log.Println("======================================================")
	log.Println(err)
	log.Println("======================================================")
	if err != nil {
		return
	}
	var resp response.EventHomeResponse
	resp.EventList = gameEvent
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	b, err := json.Marshal(resp)
	log.Println("======================================================")
	log.Println(err)
	log.Println("======================================================")
	return string(b), nil
}
