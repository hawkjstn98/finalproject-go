package repository

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/mongoConfig"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type User struct {
	id          string `json:"id" bson:"_id"`
	name        string `json:"name" bson:"name"`
	country     string `json:"country" bson:"country"`
	phoneNumber int32  `json:"phoneNumber" bson:"phoneNumber"`
}

func LoadAllUserData() (result []User) {
	dbconnect, client := mongoConfig.Configuration();
	if (dbconnect) {
		collection := client.Database("Profile").Collection("User")
		cursor, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Println("Document Error: ", err)
			return
		}

		for cursor.Next(context.Background()) {
			user := User{}
			err := cursor.Decode(&user)
			result = append(result, user)
			if err != nil {
				log.Println("Data Error", err)
			}
		}
	}

	return
}
