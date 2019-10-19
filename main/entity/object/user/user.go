package user

import "go.mongodb.org/mongo-driver/bson/primitive"

const UserCollection = "users"

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Country     string `json:"country" bson:"country"`
	PhoneNumber int32  `json:"phoneNumber" bson:"phoneNumber"`
}

