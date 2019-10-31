package user_services

import (
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
)


func GetAllUserData() string {
	user := user_repository.LoadAllUserData()
	result, _ := json.Marshal(user)
	return string(result)
}

func RegisterUser(req *request.RegisterRequest) string {
	response := new(response.RegisterResponse)
	if req.Email==""||req.Username==""||req.PhoneNumber==""||req.Password=="" {
		response.Message = "Failed to Register User, Field Cannot be Empty"
		response.ResponseCode = "FAILED"
		result, _ := json.Marshal(response)
		return string(result)
	}
	var usr = new(user.User)
	fmt.Println("Request :",req)

	usr.Email = req.Email
	usr.Password = helper.SavePassword(req.Password)
	usr.PhoneNumber = req.PhoneNumber
	usr.Username = req.Username

	fmt.Println("Data :",usr)

	res, msg := user_repository.RegisterUser(usr)


	if res{
		response.Message = "Successfully Register User"
		response.ResponseCode = "SUCCESS"
	}else{
		response.Message = "Failed to Register User, Duplicate "+msg
		response.ResponseCode = "FAILED"
	}

	result, _ := json.Marshal(response)
	return string(result)
}