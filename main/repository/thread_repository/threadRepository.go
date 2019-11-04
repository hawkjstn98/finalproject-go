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

func GetThreadPage() (result []*forum.ThreadPage) {
	cursor, err := threadCollection.Find(context.TODO(), bson.D{{}})

	if err != nil{
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()){
		var thread forum.Thread
		//log.Println("Cursor: ", cursor)
		err := cursor.Decode(&thread)
		//log.Println("Decode: ", &thread)

		//
		if err != nil {
			log.Println("Data Error", err)
			return
		}
		//log.Println("Thread: ", thread.MakerUsername)
		filter := bson.M{"username": thread.MakerUsername}
		var userThread user.User
		ctx := context.Background()
		cursorUser := userCollection.FindOne(ctx, filter)
		cursorUser.Decode(&userThread)
		log.Println("cursorUser: ", &userThread)

		var currThread forum.ThreadPage

		currThread.Id = thread.Id
		currThread.Timestamp = thread.Timestamp
		currThread.Name = thread.Name
		currThread.Category = thread.Category
		currThread.MakerUsername = thread.MakerUsername
		currThread.MakerImage = userThread.ProfileImage
		currThread.Description = thread.Description
		currThread.CommentNumber = len(thread.CommentList)

		result = append(result, &currThread)
	}



	return result

}
