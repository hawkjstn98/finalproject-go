package bookmark_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/bookmark"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var client = dbhealthcheck.Conf.MongoClient
var eventCollection = client.Database(mongo_constant.DBName).Collection(request.BookmarkCollection)

func CreateBookmark(req *bookmark.ObjectBookmark) (err error) {
	insertRes, err := eventCollection.InsertOne(context.TODO(), req)

	if err != nil {
		log.Println("CreateBookmark : ", err)
		return
	}

	log.Println("CreateBookmark : ", insertRes)
	return
}

func FindBookmark(req *bookmark.ObjectBookmark) bool {
	find := eventCollection.FindOne(context.TODO(), bson.D{{"userId", req.UserID}, {"eventId", req.EventID}})

	var res bookmark.ObjectBookmark
	err := find.Decode(&res)
	if err != nil {
		log.Println("FindBookmark : ", err)
		return false
	}

	if res.EventID != "" || res.UserID != ""{
		log.Println("FindBookmark : bookmark exists, ", res.EventID, " for user id, ", res.UserID)
		return false
	}

	log.Println("FindBookmark : Found ", res.EventID, " for user id, ", res.UserID)
	return true
}

func RemoveBookmark(req *bookmark.ObjectBookmark) (err error) {
	deleteRes, err := eventCollection.DeleteOne(context.TODO(), bson.D{{"userId", req.UserID}, {"eventId", req.EventID}})

	if err != nil {
		log.Println("RemoveBookmark : ", err)
		return
	}

	log.Println("RemoveBookmark : ", deleteRes)
	return
}