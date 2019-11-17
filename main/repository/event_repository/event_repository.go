package event_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client = dbhealthcheck.Conf.MongoClient
var eventCollection = client.Database(mongo_constant.DBName).Collection(event.EventCollection)
var userCollection = client.Database(mongo_constant.DBName).Collection(user.Collection)

func GetEventHome(page int) (result []*event.GameEvent, count int64, err error) {
	limit := int64(page * 10)
	skip := int64((page - 1) * 10)
	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"_id", 1}},
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

func CreateEvent(insert event.EventInsert) bool {
	var user user.User
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

	var gameEvent event.GameEvent
	gameEvent.MakerUsername = insert.MakerUsername
	gameEvent.Distance = 0
	gameEvent.Poster = insert.Poster
	gameEvent.Longitude = insert.Longitude
	gameEvent.Latitude = insert.Latitude
	gameEvent.DateEnd = insert.DateEnd
	gameEvent.DateStart = insert.DateStart
	gameEvent.Description = insert.Description
	gameEvent.Type = insert.Type
	gameEvent.Games = insert.Games
	gameEvent.Name = insert.Name
	gameEvent.Timestamp = time.Now()
	gameEvent.Site = insert.Site
	gameEvent.Category = insert.Category
	user.EventList = append(user.EventList, gameEvent)
	update := bson.M{"$set": bson.M{"eventList": user.EventList}}
	doc := userCollection.FindOneAndUpdate(context.TODO(), filter, update, nil)
	if doc == nil {
		log.Println("AddOrUpdate, Update Failed")
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
