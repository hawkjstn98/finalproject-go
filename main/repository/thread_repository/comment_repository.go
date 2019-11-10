package thread_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var commentCollection = client.Database(mongo_constant.DBName).Collection(forum.CommentCollection)

func GetCommentCount(id string) (int) {
	filter := bson.M{"masterThreadId": id}
	cursor, err := commentCollection.CountDocuments(context.Background(), filter)

	if err != nil {
		log.Println("Document Error: ", err)
		return 0
	}

	return int(cursor)
}

func GetCommentFromMasterID(id string) (result []*forum.ObjectComment, err error) {
	cursor, err := commentCollection.Find(context.Background(), bson.D{{"masterThreadId", id}})

	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()) {
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
