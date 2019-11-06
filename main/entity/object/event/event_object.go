package event

import "go.mongodb.org/mongo-driver/bson/primitive"

type GameEvent struct {
	Timestamp primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
	Games string `json:"games" bson:"games"`
	Description string `json:"description" bson:"description"`
}
