package thread_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/insert"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetCommentFromMasterID(id string, page int64) (result []*forum.ObjectComment, count int64, err error) {
	limit := page * 10
	skip := (page - 1) * 10
	option := &options.FindOptions{
		Skip:  &skip,
		Sort:  bson.D{{"_id", 1}},
		Limit: &limit,
	}
	cursor, err := commentCollection.Find(context.Background(), bson.D{{"masterThreadId", id}}, option)
	if err != nil {
		log.Println("Document Error: ", err)
		return
	}
	count, err = commentCollection.CountDocuments(context.Background(), bson.D{{"masterThreadId", id}})
	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	for cursor.Next(context.Background()) {
		var comment forum.ObjectComment
		err := cursor.Decode(&comment)
		if err != nil {
			log.Println("Data Error", err)
			return nil, 0, err
		}
		result = append(result, &comment)
	}

	return result, count, nil
}

func CreateThreadComment(request *insert.ThreadCommentInsert) (bool, string) {

	res, err := commentCollection.InsertOne(context.TODO(), request)

	if err != nil {
		log.Println("Error Inserting to DB: ", err)
		return false, "Failed To Create Thread Comment to DB"
	}

	log.Println("Query Result: ", res)
	return true, "Success creating Thread Comment"
}