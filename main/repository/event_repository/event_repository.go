package event_repository

import (
	"context"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func SearchEvent(page int, key string) (result []*event.GameEvent, count int64, err error)  {
	limit := int64(page * 10)
	skip := int64((page - 1) * 10)
	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"_id", 1}},
		Limit: &limit,
	}

	filter := bson.M{"name": primitive.Regex{Pattern: "^"+key, Options: "i"}}

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

func CreateEvent(insert *event.EventInsert) (bool, error) {
	var user user.User
	filter := bson.M{"username": insert.MakerUsername}

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil{
		log.Println("User doesn't exist: ", err)
		return false, err
	}

	insertRes, err := eventCollection.InsertOne(context.TODO(), insert)
	if err != nil {
		log.Println("Unable to create comment: ", err)
		return false, err
	}

	log.Println("Create Event: ", insertRes)

	//var gameEvent event.GameEvent
	//gameEvent.MakerUsername = insert.MakerUsername
	//gameEvent.Distance = 0
	//gameEvent.Poster = insert.Poster
	//gameEvent.Longitude = insert.Longitude
	//gameEvent.Latitude = insert.Latitude
	//gameEvent.DateEnd = insert.DateEnd
	//gameEvent.DateStart = insert.DateStart
	//gameEvent.Description = insert.Description
	//gameEvent.Type = insert.Type
	//gameEvent.Games = insert.Games
	//gameEvent.Name = insert.Name
	//gameEvent.Timestamp = time.Now()
	//gameEvent.Site = insert.Site
	//gameEvent.Category = insert.Category
	//var x interface{} = insertRes.InsertedID
	//gameEvent.ID = x.(primitive.ObjectID)
	//user.EventList = append(user.EventList, gameEvent)
	//update := bson.M{"$set": bson.M{"eventList": user.EventList}}
	//doc := userCollection.FindOneAndUpdate(context.TODO(), filter, update, nil)
	//if doc == nil {
	//	log.Println("AddOrUpdate, Update Failed")
	//	return false
	//}
	return true, nil
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

func MyEvent(username string) (event []event.GameEvent, message string, status bool) {
	var user user.User
	filter := bson.M{"username": username}

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		fmt.Println("User Not Found")
		return nil, "User Not Found", false
	}

	return user.EventList, "Success Find user Event", true
}
