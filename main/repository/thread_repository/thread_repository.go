package thread_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client = dbhealthcheck.Conf.MongoClient
var threadCollection = client.Database(mongo_constant.DBName).Collection(forum.ThreadCollection)
var userCollection = client.Database(mongo_constant.DBName).Collection(user.UserCollection)

func GetThreadPage(page int) (result []*forum.Thread) {
	cursor, err := threadCollection.Find(context.Background(), bson.D{{}}, options.Find().SetLimit(int64(int(page) * 10)))
	if err != nil{
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()){
		var thread forum.Thread
		err := cursor.Decode(&thread)
		if err != nil {
			log.Println("Data Error", err)
			return
		}
		result = append(result, &thread)
	}

	return result
}

func GetThreadCategory(category string) (result []*forum.Thread) {
	filter := bson.M{"category": category}
	cursor, err := threadCollection.Find(context.TODO(), filter)

	if err != nil{
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()){
		var thread forum.Thread
		err := cursor.Decode(&thread)
		if err != nil {
			log.Println("Data Error", err)
			return
		}
		result = append(result, &thread)
	}

	return result
}

func CreateThread(request *forum.Thread) (bool, string) {
	res, err := threadCollection.InsertOne(context.TODO(), request)

	if err != nil {
		log.Println(err)
		return false, "Failed To Create Thread to DB"
	}

	log.Println(res)

	return true, "Success creating Thread"
}
