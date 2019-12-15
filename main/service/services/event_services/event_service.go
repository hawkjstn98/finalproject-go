package event_services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/gcp_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/bookmark"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/bookmark_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/event_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"googlemaps.github.io/maps"
	"log"
	"math"
	"strconv"
)

func GetEventHome(req *request.EventHomeRequest) (res string, err error) {
	if req.Page < 1 {
		return "", fmt.Errorf("invalid paging")
	}
	gameEvents, count, err := event_repository.GetEventHome(req.Page)
	if err != nil {
		return
	}
	for i, gameEvent := range gameEvents{
		img, err := user_repository.GetUserImage(gameEvent.MakerUsername)
		if err != nil {
			log.Println(err)
		}
		gameEvents[i].MakerImage = img
		var objBookmark bookmark.ObjectBookmark
		objBookmark.EventID = gameEvent.ID.Hex()
		objBookmark.UserID = req.UserId
		gameEvent.BookmarkStatus = strconv.FormatBool(bookmark_repository.FindBookmark(&objBookmark))
	}
	page := float64(count / 10)
	page = math.Floor(page)
	maxPage := int64(page)
	if count%10 > 0 {
		maxPage = maxPage + 1
	}
	var resp response.EventHomeResponse
	eventList := MapToEventList(req.Latitude, req.Longitude, gameEvents)
	resp.EventList = eventList
	resp.MaxPage = maxPage
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	b, err := json.Marshal(resp)
	return string(b), nil
}

func SearchEvent(req *request.SearchEventRequest) string {
	if req.Page < 1 {
		return "Pagination Error, page cannot be 1"
	}
	gameEvents, count, err := event_repository.SearchEvent(req.Page, req.SearchKey)

	if err != nil {
		return ""
	}

	if len(gameEvents) < 0 || gameEvents == nil {
		return "No Event with this name available"
	}
	for i, gameEvent := range gameEvents{
		img, err := user_repository.GetUserImage(gameEvent.MakerUsername)
		if err != nil {
			log.Println(err)
		}
		gameEvents[i].MakerImage = img
	}
	page := float64(count / 10)
	page = math.Floor(page)
	maxPage := int64(page)
	if count%10 > 0 {
		maxPage = maxPage + 1
	}
	var resp response.EventHomeResponse
	eventList := MapToEventList(req.Latitude, req.Longitude, gameEvents)
	resp.EventList = eventList
	resp.MaxPage = maxPage
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	results, err := json.Marshal(resp)

	return string(results)
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
		latitudes  []string
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
	if req.EventId == "" || req.UserId == "" {
		return "", fmt.Errorf("invalid thread id or user id")
	}
	event, err := event_repository.GetEvent(req.EventId)
	if err != nil {
		return
	}

	if event[0].Site == "Online" {
		event[0].Latitude = req.UserLatitude
		event[0].Longitude = req.UserLongitude
	}
	var resp response.EventDetailResponse
	events := MapToEventList(req.UserLatitude, req.UserLongitude, event)
	resp.Event = events[0]
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"

	var objBookmark bookmark.ObjectBookmark
	objBookmark.EventID = req.EventId
	objBookmark.UserID = req.UserId
	events[0].BookmarkStatus = strconv.FormatBool(bookmark_repository.FindBookmark(&objBookmark))

	if req.UserLatitude == "" || req.UserLongitude == "" {
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

func MyEventService(req *request.MyEventRequest) string {
	resp := new(response.MyEventResponse)

	res, msg, status := event_repository.MyEvent(req.Username)

	fmt.Print("result : ", res)

	if !status {
		resp.Response.ResponseCode = "FAILED FETCH MY EVENT"
		resp.Response.Message = msg
		resp.EventList = nil
	}

	resp.Response.ResponseCode = "SUCCESS FETCH MY EVENT"
	resp.Response.Message = msg

	if 0 == len(res) {
		resp.Response.Message = "Empty Event"
		responses, _ := json.Marshal(resp)
		resp.EventList = nil
		return string(responses)
	}

	var latitude []string
	var longitude []string

	for i := 0; i < len(res); i++ {
		res[i].BookmarkStatus = "true"
		if res[i].Site == "Online" {
			latitude = append(latitude, req.Latitude)
			longitude = append(longitude, req.Longitude)
			continue
		}
		latitude = append(latitude, res[i].Latitude)
		longitude = append(longitude, res[i].Longitude)
	}

	dist := CountDistance(req.Latitude, req.Longitude, latitude, longitude)

	for i := 0; i < len(dist); i++ {
		res[i].Distance = dist[i]
	}
	resp.EventList = res

	responses, _ := json.Marshal(resp)
	return string(responses)
}
