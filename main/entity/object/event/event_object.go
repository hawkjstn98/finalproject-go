package event

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const EventCollection = "event"

type GameEvent struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Timestamp     time.Time          `json:"timestamp" bson:"timestamp"`
	Name          string             `json:"name" bson:"name"`
	MakerUsername string             `json:"makerusername" bson:"makerusername"`
	Type          string             `json:"type" bson:"type"`
	Games         []string           `json:"games" bson:"games"`
	Category      []string           `json:"category" bson:"category"`
	Description   string             `json:"description" bson:"description"`
	Site          string             `json:"site" bson:"site"`
	DateStart     time.Time          `json:"dateStart" bson:"dateStart"`
	DateEnd       time.Time          `json:"dateEnd" bson:"dateEnd"`
	Latitude      string             `json:"latitude" bson:"latitude"`
	Longitude     string             `json:"longitude" bson:"longitude"`
	Poster        string             `json:"poster" bson:"poster"`
	Distance      float32            `json:"distance" bson:"distance"`
}
