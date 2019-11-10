package forum_services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
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
		return "", errors.New("invalid comment paging")
	}
	commentStartIndex := (req.Page * 10) - ((req.Page - 1) * 10)
	comments, err := thread_repository.GetThreadDetail(req.ThreadID)
	if err != nil {
		return
	}
	if req.Page == 1 {
		commentStartIndex -= 10
	}
	commentEndIndex := len(comments) - commentStartIndex
	comments = comments[commentStartIndex:commentEndIndex]
	var response response.ThreadDetailResponse
	threads := MapThreadToPage(thread)
	commentsPage := MapCommentToPage(comments)
	response.Thread = &(threads[0])
	response.CommentList = commentsPage
	response.Response.Message = "SUCCESS"
	response.Response.ResponseCode = "200"
	b, err := json.Marshal(response)
	return string(b), nil
}

func MapCommentToPage(comments []*forum.ObjectComment) (commentsPage []*forum.ObjectComment) {
	for i := range comments {
		var currComment forum.ObjectComment
		imageLink := user_repository.GetUserImage(comments[i].Username)

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
