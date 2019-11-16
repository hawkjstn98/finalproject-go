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
		return "", fmt.Errorf("invalid comment paging")
	}
	comments, maxPage, err := thread_repository.GetCommentFromMasterID(req.ThreadID, req.Page)
	if err != nil {
		return
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
