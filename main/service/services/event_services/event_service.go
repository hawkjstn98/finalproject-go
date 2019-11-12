package event_services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/gcp_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/event_repository"
	"googlemaps.github.io/maps"
	"log"
)

func GetEventHome(req *request.EventHomeRequest) (res string, err error) {
	if req.Page < 1 {
		return "", fmt.Errorf("invalid paging")
	}
	gameEvent, err := event_repository.GetEventHome(req.Page)
	if err != nil {
		return
	}
	var resp response.EventHomeResponse
	eventList := MapToEventList(req, gameEvent)
	resp.EventList = eventList
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	b, err := json.Marshal(resp)
	return string(b), nil
}

func CreateEvent(req event.EventInsert) string {

	ceResponse := new(response.CreateEventResponse)

	res := event_repository.CreateEvent(req)

	if res {
		ceResponse.Response.ResponseCode = "SUCCESS"
		ceResponse.Response.Message = "Success Creating Event"
	} else {
		ceResponse.Response.ResponseCode = "FAILED"
		ceResponse.Response.Message = "Failed Creating Event, Please Contact Our Customer Support : +62895348810240"
	}

	resp, _ := json.Marshal(ceResponse)
	return string(resp)
}

func CountDistance(usrLatitude string, usrLongitude string, dataLatitude []string, dataLongitude []string) []float32 {
	count := len(dataLatitude)
	origin := make([]string, count)
	destinations := make([]string, count)
	for i := 0; i < count; i++ {
		origin[i] = usrLatitude + "|" + usrLongitude
		destinations[i] = dataLatitude[i] + "|" + dataLongitude[i]
	}
	distMatrixRequest := new(maps.DistanceMatrixRequest)
	distMatrixRequest.Origins = origin
	distMatrixRequest.Destinations = destinations

	m, err := maps.NewClient(maps.WithAPIKey(gcp_constant.APIKEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	resp, err := m.DistanceMatrix(context.Background(), distMatrixRequest)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	distances := make([]float32, count)
	element := new(maps.DistanceMatrixElement)
	for i := 0; i < count; i++ {
		element = resp.Rows[i].Elements[i]
		distances[i] = float32(element.Distance.Meters/1000)
	}

	return distances
}

func MapToEventList(req *request.EventHomeRequest, events []*event.GameEvent) (eventList []*event.GameEvent) {
	var (
		longitudes []string
		latitudes []string
	)
	for _, e := range events {
		longitudes = append(longitudes, e.Longitude)
		latitudes = append(latitudes, e.Latitude)
	}

	distances := CountDistance(req.Latitude, req.Longitude, latitudes, longitudes)
	for i, e := range events {
		e.Timestamp = e.ID.Timestamp()
		e.Distance = distances[i]
		eventList = append(eventList, e)
	}
	return
}
