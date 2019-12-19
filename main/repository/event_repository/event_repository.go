package event_repository

import (
	"context"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/bookmark"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client = dbhealthcheck.Conf.MongoClient
var eventCollection = client.Database(mongo_constant.DBName).Collection(event.EventCollection)
var userCollection = client.Database(mongo_constant.DBName).Collection(user.Collection)
var bookmarkCollection = client.Database(mongo_constant.DBName).Collection(request.BookmarkCollection)

func GetEventHome(page int) (result []*event.GameEvent, count int64, err error) {
	limit := int64(page * 10)
	skip := int64((page - 1) * 10)
	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"dateStart", -1}},
		Limit: &limit,
	}
	cursor, err := eventCollection.Find(context.TODO(), bson.D{{}}, option)
	if err != nil {
		log.Println("GetEventHome : Document Error, ", err)
		return
	}
	count, err = eventCollection.CountDocuments(context.Background(), bson.D{{}})
	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()) {
		var gameEvent event.GameEvent
		err := cursor.Decode(&gameEvent)
		if err != nil {
			log.Println("GetEventHome : Decode Error, ", err)
			return nil, 0, err
		}
		result = append(result, &gameEvent)
	}

	return result, count, nil
}

func SearchEvent(page int, key string) (result []*event.GameEvent, count int64, err error)  {
	limit := int64(page * 10)
	skip := int64((page - 1) * 10)
	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"_id", 1}},
		Limit: &limit,
	}

	filter := bson.M{"name": primitive.Regex{Pattern: "("+key+")", Options: "i"}}

	cursor, err := eventCollection.Find(context.TODO(), filter, option)

	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	count, err = eventCollection.CountDocuments(context.Background(), filter)

	if err != nil {
		log.Println("Document Error: ", err)
		return
	}


	for cursor.Next(context.Background()) {
		var gameEvent event.GameEvent
		err := cursor.Decode(&gameEvent)
		if err != nil {
			log.Println("GetEventHome : Decode Error, ", err)
			return nil, 0, err
		}
		result = append(result, &gameEvent)
	}

	return result, count, nil
}

func CreateEvent(insert event.EventInsert) bool {
	var user user.UserWithId
	filter := bson.M{"username": insert.MakerUsername}

	err1 := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err1 != nil{
		return false
	}

	insertRes, err := eventCollection.InsertOne(context.TODO(), insert)

	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("event : ", insertRes)

	var bookmarkObject bookmark.ObjectBookmark
	bookmarkObject.EventID = insertRes.InsertedID.(primitive.ObjectID).Hex()
	bookmarkObject.UserID = user.Id.Hex()
	_, err = bookmarkCollection.InsertOne(context.TODO(), bookmarkObject)
	if err!= nil {
		log.Println("Insert bookmark failed")
		return false
	}
	return true
}

func GetEvent(id string) (result []*event.GameEvent, err error) {
	Id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": Id}
	cursor, err := eventCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	if cursor.Next(context.Background()) {
		var event event.GameEvent
		err := cursor.Decode(&event)
		if err != nil {
			result = nil
			log.Println("Data Error", err)
			return nil, err
		}
		result = append(result, &event)
	}
	return result, nil
}

func MyEvent(username string) (events []*event.GameEvent, message string, status bool) {
	var user user.UserWithId
	filter := bson.M{"username": username}

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		fmt.Println("User Not Found")
		return nil, "User Not Found", false
	}

	id := user.Id.Hex()
	filters := bson.M{"userId": id}
	cursor, err := bookmarkCollection.Find(context.TODO(), filters)

	var eventId []string
	for cursor.Next(context.Background()) {
		var bookmarks bookmark.ObjectBookmark
		err := cursor.Decode(&bookmarks)
		if err != nil {
			eventId = nil
			log.Println("Data Error", err)
			return nil, "ERROR", false
		}
		eventId = append(eventId, (&bookmarks).EventID)
	}

	var eventList []*event.GameEvent
	for i := 0; i < len(eventId); i++ {
		var evnt event.GameEvent
		eventid, _ := primitive.ObjectIDFromHex(eventId[i])
		filterEvent := bson.M{"_id": eventid}
		err := eventCollection.FindOne(context.TODO(), filterEvent).Decode(&evnt)
		fmt.Println("Error: ", err)
		if err!= nil {
			log.Println("Insert bookmark failed")
			return nil, "Data Error", false
		}
		eventList = append(eventList, &evnt)
	}
	return eventList, "Success Find user Event", true
}
