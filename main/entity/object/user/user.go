package user

import "go.mongodb.org/mongo-driver/bson/primitive"

const UserCollection = "users"

type User struct {
	Id         	 	primitive.ObjectID `json:"id" bson:"_id"`
	Name        	string `json:"name" bson:"name"`
	PhoneNumber 	int32  `json:"phoneNumber" bson:"phoneNumber"`
	Timestamp  	 	primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Email 			string `json:"email" bson:"email"`
	Username 		string `json:"username" bson:"username"`
	GameList    	[]string `json:"gameList" bson:"gameList"`
	EventList   	[]string `json:"eventList" bson:"eventList"`
	ProfileImage 	string `json:"profileImage" bson:"profileImage"`
}

