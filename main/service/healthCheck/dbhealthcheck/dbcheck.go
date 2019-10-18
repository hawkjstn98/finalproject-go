package dbhealthcheck

import (
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/mongoConfig"
)

func CheckAndGetConfig() mongoConfig.MongoConfig{
	mongoConfig := mongoConfig.Configuration()

	if mongoConfig.MongoClient == nil || mongoConfig.Status == false {
		fmt.Println("Connection to Db Failed")
		return mongoConfig
	}

	return mongoConfig
}
