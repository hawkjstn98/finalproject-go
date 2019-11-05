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
		filter := bson.M{"email": usr.Email}
		err := userCollection.FindOne(context.TODO(), filter).Decode(&userEmail)
		if err != nil {
			log.Print(err)
		}

		if userEmail.Email == "" {
			insertRes, err := userCollection.InsertOne(context.TODO(), usr)
			if err != nil {
				log.Print(err)
			}
			log.Println("user: ", insertRes)
			return true, ""
		}

		return false, "Email"
	}

	return false, "Username"
}

func UserLogin(email string, password string) (bool, string) {

	var user user.User

	filter := bson.M{
		"$and": []bson.M{
			bson.M{"email": email},
			bson.M{"password": password},
		},
	}

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return false, "User Not Found"
	}

	return true, user.Username
}

func GetUserImage(username string) string {
	filter := bson.M{"username": username}
	var userThread user.User
	ctx := context.Background()
	cursorUser := userCollection.FindOne(ctx, filter)
	cursorUser.Decode(&userThread)
	log.Println("cursorUser: ", &userThread)
	return userThread.ProfileImage
}

func AddOrUpdateGameList(username string, gameList [] string) (bool, string, interface{}) {
	filter := bson.M{"username": username}

	update := bson.M{"$set": bson.M{"gameList": gameList}}

	doc := userCollection.FindOneAndUpdate(context.TODO(), filter, update, nil)

	if doc == nil {
		log.Println("AddOrUpdate, Update Failed")
		return false, "User Not Found", doc
	}

	return true, "Update Success", doc
}
