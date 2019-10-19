package user_repository

import (
	"context"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)


func LoadAllUserData(client *mongo.Client) (result []*user.User) {

		collection := client.Database(mongo_constant.DBName).Collection(user.UserCollection)
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cursor, err := collection.Find(ctx, bson.D{{}})

		if err != nil {
			log.Println("Document Error: ", err)
			return
		}

		defer cursor.Close(ctx)

		fmt.Println("print: ", collection.Name())

		for cursor.Next(context.Background()) {
			var user user.User
			err := cursor.Decode(&user)
			fmt.Println("Document: ", &user)
			result = append(result, &user)
			if err != nil {
				log.Println("Data Error", err)
			}
		}

	return result
}
