package event_services

import (
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/bookmark"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/bookmark_repository"
)

func Bookmark(req *request.BookmarkRequest) (res string, err error) {
	if req.UserID == "" || req.EventID == "" {
		return "", fmt.Errorf("invalid request")
	}

	var objBookmark bookmark.ObjectBookmark
	objBookmark.EventID = req.EventID
	objBookmark.UserID = req.UserID

	double := bookmark_repository.FindBookmark(&objBookmark)
	if double{
		return
	}
	err = bookmark_repository.CreateBookmark(&objBookmark)
	if err != nil {
		return
	}
	var resp response.BookmarkResponse
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	b, err := json.Marshal(resp)
	return string(b), nil
}

func RemoveBookmark(req *request.BookmarkRequest) (res string, err error) {
	if req.UserID == "" || req.EventID == "" {
		return "", fmt.Errorf("invalid request")
	}

	var objBookmark bookmark.ObjectBookmark
	objBookmark.EventID = req.EventID
	objBookmark.UserID = req.UserID

	err = bookmark_repository.RemoveBookmark(&objBookmark)
	if err != nil {
		return
	}
	var resp response.BookmarkResponse
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	b, err := json.Marshal(resp)
	return string(b), nil
}