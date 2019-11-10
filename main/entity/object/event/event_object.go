package event

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const EventCollection = "event"

type GameEvent struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Name string `json:"name" bson:"name"`
	MakerUsername string `json:"makerusername" bson:"makerusername"`
	Type string `json:"type" bson:"type"`
	Games []string `json:"games" bson:"games"`
	Category []string `json:"category" bson:"category"`
	Description string `json:"description" bson:"description"`
	Site string `json:"site" bson:"site"`
	DateStart int32 `json:"dateStart" bson:"dateStart"`
	DateEnd int32 `json:"dateEnd" bson:"dateEnd"`
	Location string `json:"location" bson:"location"`
	Poster string `json:"poster" bson:"poster"`
}
