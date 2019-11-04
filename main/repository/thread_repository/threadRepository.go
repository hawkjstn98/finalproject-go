package thread_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var client = dbhealthcheck.Conf.MongoClient
var threadCollection = client.Database(mongo_constant.DBName).Collection(forum.ThreadCollection)
var userCollection = client.Database(mongo_constant.DBName).Collection(user.UserCollection)

func GetThreadPage() (result []*forum.Thread) {
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
			return
		}
		result = append(result, &thread)
	}

	return result

}
