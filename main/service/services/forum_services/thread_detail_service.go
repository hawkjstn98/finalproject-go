package forum_services

import (
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/insert"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"log"
	"math"
)

func GetThreadDetail(req *request.ThreadDetailRequest) (res string, err error) {
	if req.ThreadID == "" {
		return "", fmt.Errorf("invalid thread id")
	}
	thread, err := thread_repository.GetThread(req.ThreadID)
	if err != nil {
		return
	}
	if req.Page < 1 {
		return "", fmt.Errorf("invalid comment paging")
	}
	comments, count, err := thread_repository.GetCommentFromMasterID(req.ThreadID, req.Page)
	if err != nil {
		return
	}
	page := float64(count / 10)
	page = math.Floor(page)
	maxPage := int64(page)
	if count % 10 > 0 {
		maxPage = maxPage + 1
	}
	var resp response.ThreadDetailResponse
	threads := MapThreadToPage(thread)
	commentsPage := MapCommentToPage(comments)
	resp.Thread = threads[0]
	resp.CommentList = commentsPage
	resp.MaxPage = maxPage
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	b, err := json.Marshal(resp)
	return string(b), nil
}

func MapCommentToPage(comments []*forum.ObjectComment) (commentsPage []*forum.ObjectComment) {
	for i := range comments {
		var currComment forum.ObjectComment
		imageLink, err := user_repository.GetUserImage(comments[i].Username)
		if err != nil{
			log.Println(err)
		}
		currComment.Id = comments[i].Id
		currComment.ThreadMasterID = comments[i].ThreadMasterID
		currComment.Timestamp = comments[i].Id.Timestamp()
		currComment.Username = comments[i].Username
		currComment.ProfileImage = imageLink
		currComment.ThreadComment = comments[i].ThreadComment

		commentsPage = append(commentsPage, &currComment)
	}
	return commentsPage
}

func CreateThreadComment(threadRequest *request.CreateThreadCommentRequest) string {
	res := new(response.CreateThreadCommentResponse)

	if "" == threadRequest.MasterThreadID || "" == threadRequest.MakerUsername || "" == threadRequest.Category || "" == threadRequest.Description || threadRequest.Timestamp.IsZero() {
		res.Response.Message = "Invalid Request Format"
		res.Response.ResponseCode = "Failed To Create Comment"
		result, err := json.Marshal(res)
		if err != nil{
			log.Println("Error Marshal ", err)
		}
		return string(result)
	}

	var thread = new(insert.ThreadCommentInsert)
	thread.MasterThreadID = threadRequest.MasterThreadID
	thread.Timestamp = threadRequest.Timestamp
	thread.Description = threadRequest.Description
	thread.Category = threadRequest.Category
	thread.MakerUsername = threadRequest.MakerUsername

	resp, msg := thread_repository.CreateThreadComment(thread)

	if resp {
		res.Response.Message = msg
		res.Response.ResponseCode = "Create Thread Comment Success"
	} else {
		res.Response.Message = "Create Thread Comment failed, " + msg
		res.Response.ResponseCode = "Create Thread Comment Failed"
	}

	result, err := json.Marshal(res)
	if err != nil{
		log.Println("Error Marshal ", err)
	}
	return string(result)
}