package thread_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/insert"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
)

var client = dbhealthcheck.Conf.MongoClient
var threadCollection = client.Database(mongo_constant.DBName).Collection(forum.ThreadCollection)

func GetThreadPage(page int) (result []*forum.Thread, err error) {
	var limit, skip int64
	if(page != 0){
		limit = int64(page * 10)
		skip = int64((page - 1) * 10)
	} else {
		return
	}
	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"_id", 1}},
		Limit: &limit,
	}
	cursor, err := threadCollection.Find(context.Background(), bson.D{{}}, option)
	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()) {
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

func GetThreadCategory(category *request.ThreadCategoryRequest) (result []*forum.Thread) {
	filter := bson.M{"category": category.Category}
	var limit, skip int64
	if(category.Page != 0){
		limit = int64(category.Page * 10)
		skip = int64((category.Page - 1) * 10)
	} else{
		return
	}

	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"_id", 1}},
		Limit: &limit,
	}
	cursor, err := threadCollection.Find(context.TODO(), filter, option)
	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()) {
		var thread forum.Thread
		err := cursor.Decode(&thread)
		if err != nil {
			log.Println("Data Error", err)
			return
		}
		result = append(result, &thread)
	}
	log.Println(result)

	return result
}

func GetThread(id string) (result []*forum.Thread, err error) {
	Id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": Id}
	cursor, err := threadCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	if cursor.Next(context.Background()) {
		var thread forum.Thread
		err := cursor.Decode(&thread)
		if err != nil {
			result = nil
			log.Println("Data Error", err)
			return nil, err
		}
		result = append(result, &thread)
	}
	return result, nil
}

func CreateThread(request *insert.ThreadInsert) (bool, string) {

	res, err := threadCollection.InsertOne(context.TODO(), request)

	if err != nil {
		log.Println(err)
		return false, "Failed To Create Thread to DB"
	}

	log.Println(res)

	return true, "Success creating Thread"
}

func GetThreadCount(category string) (int){
	if category == "" {
		cursor, err := threadCollection.CountDocuments(context.Background(), bson.D{{}})
		if err != nil {
			log.Println("Document Error: ", err)
			return 0
		}
		return int(cursor)
	} else {
		filter := bson.M{"category": category}
		cursor, err := threadCollection.CountDocuments(context.Background(), filter)
		if err != nil {
			log.Println("Document Error: ", err)
			return 0
		}
		return int(cursor)
	}



}
