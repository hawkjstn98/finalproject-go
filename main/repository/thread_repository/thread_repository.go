package thread_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
)

var client = dbhealthcheck.Conf.MongoClient
var threadCollection = client.Database(mongo_constant.DBName).Collection(forum.ThreadCollection)

func GetThreadPage(page int) (result []*forum.Thread, err error) {
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
			return nil, err
		}
		result = append(result, &thread)
	}

	return result, nil
}

func GetThreadCategory(category *request.ThreadCategoryRequest) (result []*forum.Thread) {
	filter := bson.M{"category": category.Category}
	cursor, err := threadCollection.Find(context.TODO(), filter, options.Find().SetLimit(int64(int(category.Page) * 10)))

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

func GetThread(id string) (result *forum.Thread, err error){
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

func GetThreadDetail(id string) (result []*forum.ObjectComment, err error){
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
