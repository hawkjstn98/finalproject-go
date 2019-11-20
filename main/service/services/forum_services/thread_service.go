package forum_services

import (
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/insert"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	//"log"
)

func GetThreadPage(page *request.ThreadRequest) string {
	threads, _ := thread_repository.GetThreadPage(page.Page)
	threadsPage := MapThreadToPage(threads)

	var resp response.ThreadResponse
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	resp.Thread = threadsPage
	resp.MaxPage = GetMaxPage("home", "")

	result, _ := json.Marshal(resp)
	return string(result)
}

func GetSearchPage(request *request.SearchThreadRequest) string {
	threads, _ := thread_repository.GetSearchPage(request.Page, request.SearchKey)
	threadWithPage := MapThreadToPage(threads)

	var resp response.SearchThreadResponse
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	resp.Thread = threadWithPage
	resp.MaxPage = GetMaxPage("", request.SearchKey)

	result, _ := json.Marshal(resp)
	return string(result)
}

func GetThreadCategoryPage(category *request.ThreadCategoryRequest) string {
	threads := thread_repository.GetThreadCategory(category)
	threadsPage := MapThreadToPage(threads)

	var resp response.ThreadResponse
	resp.Response.Message = "SUCCESS"
	resp.Response.ResponseCode = "200"
	resp.Thread = threadsPage
	resp.MaxPage = GetMaxPage(category.Category, "")

	result, _ := json.Marshal(resp)
	return string(result)
}

func MapThreadToPage(threads []*forum.Thread) (threadsPage []*forum.Thread) {
	for i := range threads {
		var currThread forum.Thread
		imageLink := user_repository.GetUserImage(threads[i].MakerUsername)

		currThread.Id = threads[i].Id
		currThread.Timestamp = threads[i].Id.Timestamp()
		currThread.Name = threads[i].Name
		currThread.Category = threads[i].Category
		currThread.MakerUsername = threads[i].MakerUsername
		currThread.MakerImage = imageLink
		currThread.Description = threads[i].Description
		currThread.CommentCount = thread_repository.GetCommentCount(threads[i].Id.Hex())

		threadsPage = append(threadsPage, &currThread)
	}
	return threadsPage
}

func GetMaxPage(category string, key string) int {
	var threadCount int
	if key != "" {
		threadCount = thread_repository.GetSearchCount(key)
	} else {
		threadCount = thread_repository.GetThreadCount(category)
	}
	if threadCount % 10 == 0 {
		return threadCount/10
	} else {
		return threadCount/10 + 1
	}
}

func GetStart(end int) (int) {
	if int(end) > 10 {
		if int(end)%10 == 0 {
			return int(end) - 10
		} else {
			return (int(end) / 10) * 10
		}
	} else {
		return 0
	}
}

func CreateThread(threadRequest *request.CreateThreadRequest) string {
	response := new(response.CreateThreadResponse)

	if "" == threadRequest.MakerUsername || "" == threadRequest.Name || "" == threadRequest.Category || "" == threadRequest.Description || threadRequest.Timestamp.IsZero() {
		response.Response.Message = "Invalid Request Format"
		response.Response.ResponseCode = "Failed To Add Or Update PhoneNumber"
	}

	var thread = new(insert.ThreadInsert)
	thread.Timestamp = threadRequest.Timestamp
	thread.Description = threadRequest.Description
	thread.Category = threadRequest.Category
	thread.MakerUsername = threadRequest.MakerUsername
	thread.Name = threadRequest.Name

	res, msg := thread_repository.CreateThread(thread)

	if res {
		response.Response.Message = msg
		response.Response.ResponseCode = "Create Thread Success"
	} else {
		response.Response.Message = "Create Thread failed, " + msg
		response.Response.ResponseCode = "Create Thread Failed"
	}

	result, _ := json.Marshal(response)
	return string(result)

}
