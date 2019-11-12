package event_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client = dbhealthcheck.Conf.MongoClient
var eventCollection = client.Database(mongo_constant.DBName).Collection(event.EventCollection)

func GetEventHome(page int) (result []*event.GameEvent, err error) {
	limit := int64(page * 10)
	skip := int64((page - 1) * 10)
	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"_id", 1}},
		Limit: &limit,
	}
	cursor, err := eventCollection.Find(context.TODO(), bson.D{{}}, option)

	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()) {
		var gameEvent event.GameEvent
		err := cursor.Decode(&gameEvent)
		if err != nil {
			log.Println("Data Error", err)
			return nil, err
		}
		result = append(result, &gameEvent)
	}

	return result, nil
}

func CreateEvent(insert event.EventInsert) bool {
	insertRes, err := eventCollection.InsertOne(context.TODO(), insert)

	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("event : ", insertRes)
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
