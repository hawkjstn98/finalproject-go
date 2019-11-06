package thread_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var client = dbhealthcheck.Conf.MongoClient
var threadCollection = client.Database(mongo_constant.DBName).Collection(forum.ThreadCollection)

func GetThreadPage() (result []*forum.Thread, err error) {
	cursor, err := threadCollection.Find(context.TODO(), bson.D{{}})

	if err != nil{
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()){
		var thread forum.Thread
		err := cursor.Decode(&thread)
		if err != nil {
			log.Println("Data Error", err)
			return nil, err
		}
		result = append(result, &thread)
	}

	return result, nil
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

func GetThread(id int) (result *forum.Thread, err error){
	cursor, err := threadCollection.Find(context.Background(), bson.D{{"_id", id}})

	if err != nil{
		log.Println("Document Error: ", err)
		return
	}

	if cursor.Next(context.Background()){
		err := cursor.Decode(&result)
		if err != nil {
			result = nil
			log.Println("Data Error", err)
			return nil, err
		}
	}
	return result, nil
}

func GetThreadDetail(id int) (result []*forum.ObjectComment, err error){
	cursor, err := threadCollection.Find(context.Background(), bson.D{{}})

	if err != nil{
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()){
		var comment forum.ObjectComment
		err := cursor.Decode(&comment)
		if err != nil {
			log.Println("Data Error", err)
			return nil, err
		}
		result = append(result, &comment)
	}

	return result, nil
}
