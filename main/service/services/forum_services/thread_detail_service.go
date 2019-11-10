package forum_services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
)

func GetThreadDetail(req *request.ThreadDetailRequest) (res string, err error){
	if req.ThreadID == ""{
		return "", fmt.Errorf("invalid thread id")
	}
	thread, err := thread_repository.GetThread(req.ThreadID)
	if err != nil{
		return
	}
	if req.Page < 1 {
		return "", errors.New("invalid comment paging")
	}
	commentStartIndex := (req.Page * 10) - ((req.Page - 1) * 10)
	comments, err := thread_repository.GetThreadDetail(req.ThreadID)
	if err != nil{
		return
	}
	commentEndIndex := len(comments) - commentStartIndex
	if req.Page == 1 {
		commentStartIndex -= 10
	}
	comments = comments[commentStartIndex : commentEndIndex]
	var response response.ThreadDetailResponse
	response.Thread = thread
	response.CommentList = comments
	response.Response.Message = "SUCCESS"
	response.Response.ResponseCode = "200"
	b, err := json.Marshal(response)
	//log.Println(string(b))
	return string(b), nil
}
