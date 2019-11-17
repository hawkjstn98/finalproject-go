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
	"math"
)

func GetEventHome(req *request.EventHomeRequest) (res string, err error) {
	if req.Page < 1 {
		return "", fmt.Errorf("invalid paging")
	}
	gameEvent, count, err := event_repository.GetEventHome(req.Page)
	if err != nil {
		return
	}
	page := float64(count / 10)
	page = math.Floor(page)
	maxPage := int64(page)
	if count % 10 > 0 {
		maxPage = maxPage + 1
	}
	var resp response.EventHomeResponse
	eventList := MapToEventList(req.Latitude, req.Longitude, gameEvent)
	resp.EventList = eventList
	resp.MaxPage = maxPage
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
		origin[i] = usrLatitude + "," + usrLongitude
		destinations[i] = dataLatitude[i] + "," + dataLongitude[i]
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
		distances[i] = float32(element.Distance.Meters) / 1000
	}

	return distances
}

func MapToEventList(latitude string, longitude string, events []*event.GameEvent) (eventList []*event.GameEvent) {
	var (
		longitudes []string
		latitudes []string
	)
	for _, e := range events {
		longitudes = append(longitudes, e.Longitude)
		latitudes = append(latitudes, e.Latitude)
	}

	distances := CountDistance(latitude, longitude, latitudes, longitudes)
	for i, e := range events {
		e.Timestamp = e.ID.Timestamp()
		e.Distance = distances[i]
		eventList = append(eventList, e)
	}
	return
}

func EventDetail(req *request.EventDetailRequest) (res string, err error) {
	if req.EventId == "" {
		return "", fmt.Errorf("invalid thread id")
	}
	event, err := event_repository.GetEvent(req.EventId)
	if err != nil {
		return
	}

	var resp response.EventDetailResponse
	events := MapToEventList(req.UserLatitude, req.UserLongitude, event)
	resp.Event = events[0]
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"

	if(req.UserLatitude == "" || req.UserLongitude == ""){
		resp.Distance = -1
	} else {
		var latitude, longitude []string
		latitude = append(latitude, events[0].Latitude)
		longitude = append(longitude, events[0].Longitude)
		resp.Distance = CountDistance(req.UserLatitude, req.UserLongitude, latitude, longitude)[0]
	}
	b, _ := json.Marshal(resp)
	return string(b), nil
}
