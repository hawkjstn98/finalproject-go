package user

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"time"
)

const Collection = "user"

type User struct {
	PhoneNumber 	string               `json:"phoneNumber" bson:"phoneNumber"`
	Timestamp  	 	time.Time			 `json:"timestamp" bson:"timestamp"`
	Email 			string             `json:"email" bson:"email"`
	Username 		string              `json:"username" bson:"username"`
	Password        string              `json:"password" bson:"password"`
	GameList    	[]string            `json:"gameList" bson:"gameList"`
	EventList   	[]event.GameEvent      `json:"eventList" bson:"eventList"`
	ProfileImage 	string              `json:"profileImage" bson:"profileImage"`
}

