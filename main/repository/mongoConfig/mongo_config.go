package mongoConfig

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Configuration() (status bool, client *mongo.Client)  {
	//Set Mongo Db Client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//Connect To Mongo Db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println(err)
		return false, nil
	}

	//Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
		return false, nil
	}

	return true, client
}
