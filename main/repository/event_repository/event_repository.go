package event_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
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
