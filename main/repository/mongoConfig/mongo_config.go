package mongoConfig

import (
	"context"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/constant/mongo_constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoConfig struct {
	Status bool
	MongoClient *mongo.Client
}

func Configuration() (mongoConfig MongoConfig)  {
	//Set Mongo Db Client
	clientOptions := options.Client().ApplyURI(mongo_constant.MongoDbHost)

	//Connect To Mongo Db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println(err)
		return MongoConfig{false, nil}
	}

	//Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
		return MongoConfig{false, nil}
	}

	return MongoConfig{true, client}
}
