package dbhealthcheck

import (
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper/mongoConfig"
)

//this shit must be moved somewhere but can't on entity
var Conf = CheckAndGetConfig();

func CheckAndGetConfig() mongoConfig.MongoConfig {
	mongoConfig := mongoConfig.Configuration()

	if mongoConfig.MongoClient == nil || mongoConfig.Status == false {
		fmt.Println("Connection to Db Failed")
		return mongoConfig
	}

	return mongoConfig
}
