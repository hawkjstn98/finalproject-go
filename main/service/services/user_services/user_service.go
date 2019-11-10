package user_services

import (
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/user"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/helper"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"log"
	"time"
)

func GetAllUserData() string {
	user := user_repository.LoadAllUserData()
	result, _ := json.Marshal(user)
	return string(result)
}

func RegisterUser(req *request.RegisterRequest) string {
	response := new(response.RegisterResponse)
	if "" == req.Email || "" == req.Username || "" == req.PhoneNumber || "" == req.Password {
		response.Response.Message = "Failed to Register User, Field Cannot be Empty"
		response.Response.ResponseCode = "FAILED"
		result, _ := json.Marshal(response)
		return string(result)
	}
	var usr = new(user.User)
	fmt.Println("Request :", req)

	usr.Email = req.Email
	usr.Password = helper.SavePassword(req.Password)
	usr.PhoneNumber = req.PhoneNumber
	usr.Username = req.Username
	usr.Timestamp = time.Now()

	fmt.Println("Data :", usr)

	res, msg := user_repository.RegisterUser(usr)

	if res {
		response.Response.Message = "Successfully Register User"
		response.Response.ResponseCode = "SUCCESS"
	} else {
		response.Response.Message = "Failed to Register User, Duplicate " + msg
		response.Response.ResponseCode = "FAILED"
	}

	result, _ := json.Marshal(response)
	return string(result)
}

func LoginUser(req *request.LoginRequest) string {
	response := new(response.LoginResponse)

	if "" == req.Password || "" == req.Email {
		response.Response.Message = "Invalid Username or Password Format"
		response.Response.ResponseCode = "FAILED TO LOGIN"
	}

	pass := helper.SavePassword(req.Password)

	res, msg := user_repository.UserLogin(req.Email, pass)
	fmt.Println(res)

	if res {
		response.Username = msg
		response.Response.Message = "Success Login, Welcome " + response.Username
		response.Response.ResponseCode = "Success Login"
	} else {
		response.Username = msg
		response.Response.Message = "Login Failed"
		response.Response.ResponseCode = "Email or password not found"
	}

	result, _ := json.Marshal(response)
	return string(result)
}

func AddOrUpdateGameList(req *request.AddOrUpdateGameListRequest) string {
	response := new (response.AddOrUpdateGameListResponse)

	if ""==req.Username || len(req.GameList)<=0 {
		response.Response.Message = "Invalid Request Format"
		response.Response.ResponseCode = "Failed To Insert GameList"
	}

	res, msg, results := user_repository.AddOrUpdateGameList(req.Username, req.GameList)
	log.Println(results)

	if res {
		response.Response.Message = "Successfully Add Or Update your game"
		response.Response.ResponseCode = "Update Success"
	} else {
		response.Response.Message = "Login Failed, "+msg
		response.Response.ResponseCode = "Update Failed"
	}

	result, _ := json.Marshal(response)
	return string(result)
}

func AddOrUpdatePhone(req *request.AddOrUpdatePhoneRequest) string {
	response := new(response.AddOrUpdatePhoneResponse)

	if ""==req.Username || ""==req.PhoneNumber {
		response.Response.Message = "Invalid Request Format"
		response.Response.ResponseCode = "Failed To Add Or Update PhoneNumber"
	}

	res, msg, results := user_repository.AddOrUpdatePhoneNumber(req.Username, req.PhoneNumber)
	log.Println(results)

	if res {
		response.Response.Message = "Successfully Add Or Update your PhoneNumber"
		response.Response.ResponseCode = "Update Success"
	} else {
		response.Response.Message = "Login Failed, "+msg
		response.Response.ResponseCode = "Update Failed"
	}

	result, _ := json.Marshal(response)
	return string(result)
}

func GetUserData(username string) string {

	response := new(response.UserDataResponse)

	res, usr := user_repository.GetUserData(username)

	//fmt.Println("Test",usr)

	if res {
		usr.Response.Message = "Successfully Fetch User Data"
		usr.Response.ResponseCode = "Get User Success"
	} else {
		usr.Response.Message = "Fetch User Data Failed"
		usr.Response.ResponseCode = "Get User Failed"
	}

	response = usr

	result, _ := json.Marshal(response)

	return string(result)
}
