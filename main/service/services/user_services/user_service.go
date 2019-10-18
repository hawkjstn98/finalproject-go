package user_services

import (
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/service/healthCheck/dbhealthcheck"
)

func GetAllUserData() string {
	conf := dbhealthcheck.CheckAndGetConfig()
	if conf.Status != false {
		user := user_repository.LoadAllUserData(conf.MongoClient)
		result, _ := json.Marshal(user)
		return string(result)
	}
	return "ERROR"
}