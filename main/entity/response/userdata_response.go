package response

import (
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/event"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDataResponse struct {
	Response     BasicResponse
	UserId       primitive.ObjectID `json:"id" bson:"_id"`
	PhoneNumber  string             `json:"phoneNumber" bson:"phoneNumber"`
	Email        string             `json:"email" bson:"email"`
	Username     string             `json:"username" bson:"username"`
	GameList     []string           `json:"gameList" bson:"gameList"`
	EventList    []event.GameEvent  `json:"eventList" bson:"eventList"`
	ProfileImage string             `json:"profileImage" bson:"profileImage"`
}
