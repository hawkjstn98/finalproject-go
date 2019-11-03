package user_repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/dbhealthcheck"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

var client = dbhealthcheck.Conf.MongoClient
var userCollection = client.Database(mongo_constant.DBName).Collection(user.UserCollection)

func LoadAllUserData() (result []*user.User) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cursor, err := userCollection.Find(ctx, bson.D{{}})

	if err != nil {
		log.Println("Document Error: ", err)
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(context.Background()) {
		var user user.User
		err := cursor.Decode(&user)
		result = append(result, &user)
		if err != nil {
			log.Println("Data Error", err)
		}
	}

	return result
}

func RegisterUser(usr *user.User) (bool, string) {
	var userName user.User

	filter := bson.M{"username": usr.Username}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&userName)
	if err != nil {
		log.Print(err)
	}

	if userName.Username == "" {
		var userEmail user.User
		filter := bson.M{"email":usr.Email}
		err := userCollection.FindOne(context.TODO(), filter).Decode(&userEmail)
		if err != nil {
			log.Print(err)
		}

		if userEmail.Email == "" {
			insertRes, err := userCollection.InsertOne(context.TODO(), usr)
			if err != nil {
				log.Print(err)
			}
			log.Println("user: ",insertRes)
			return true, ""
		}

		return false, "Email"
	}

	return false, "Username"
}
