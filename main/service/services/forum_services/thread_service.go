package forum_services

import (
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/insert"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"strconv"
	"log"
)

func GetThreadPage(page *request.ThreadRequest) string {
	threads, _ := thread_repository.GetThreadPage(page.Page)
	threadsPage := MapThreadToPage(threads)

	end := len(threadsPage)
	start := GetStart(end)

	result, _ := json.Marshal(threadsPage[int(start):end])
	return string(result)
}

func GetThreadCategoryPage(category *request.ThreadCategoryRequest) string {
	threads := thread_repository.GetThreadCategory(category)
	threadsPage := MapThreadToPage(threads)
	log.Println(threadsPage)

	result, _ := json.Marshal(threadsPage)
	return string(result)
}

func MapThreadToPage(threads []*forum.Thread) (threadsPage []forum.Thread) {
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

		threadsPage = append(threadsPage, currThread)
	}
	return threadsPage
}

func GetMaxPage(category *request.ThreadCategoryRequest) string {
	category.Page = 0
	var threads []*forum.Thread

	if "" == category.Category {
		threads, _ = thread_repository.GetThreadPage(category.Page)
	} else {
		threads = thread_repository.GetThreadCategory(category)
	}
	threadsPage := MapThreadToPage(threads)
	if len(threadsPage) % 10 == 0 {
		return strconv.Itoa(len(threadsPage)/10)
	} else {
		return strconv.Itoa((len(threadsPage)/10) + 1)
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
